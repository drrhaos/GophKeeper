// Package pg реализует взаимодействие с базой данной Postgres.
package pg

import (
	"context"
	"errors"
	"time"

	"gophkeeper/internal/logger"
	"gophkeeper/internal/store"
	"gophkeeper/pkg/proto"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const sourceMigrations = "file://db/migrations"

// Database хранит пул коннектов.
type Database struct {
	Conn *pgxpool.Pool
}

// NewDatabase создает новое подключение к базе данных.
func NewDatabase(uri string) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		logger.Log.Panic("Ошибка при парсинге конфигурации:", zap.Error(err))
		return nil
	}
	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.Log.Panic("Не удалось подключиться к базе данных")
		return nil
	}
	db := &Database{Conn: conn}
	return db
}

// Migrations миграция базы данных.
func Migrations(uri string) {
	m, err := migrate.New(
		sourceMigrations,
		uri)
	if err != nil {
		logger.Log.Panic("Не удалось подключиться к базе данных", zap.Error(err))
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Log.Panic("Не удалось выполнить миграцию", zap.Error(err))
	}
	logger.Log.Info("Миграции успешно применены")
}

// UserRegister добавляет нового пользоватяля в базу данных.
func (db *Database) UserRegister(ctx context.Context, login string, password string) error {
	var countRow int64
	err := db.Conn.QueryRow(ctx, `SELECT COUNT(login) FROM users WHERE login = $1`, login).Scan(&countRow)
	if err != nil {
		logger.Log.Warn("Ошибка выполнения запроса ", zap.Error(err))
		return err
	}

	if countRow != 0 {
		logger.Log.Warn("Пользователь существует")
		return store.ErrLoginDuplicate
	}

	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Warn("Ошибка при хешировании пароля ", zap.Error(err))
		return err
	}

	_, err = db.Conn.Exec(ctx,
		`INSERT INTO users (login, password, registered_at, last_time) VALUES ($1, $2, $3, $4)`, login, string(hashedPassword), time.Now(), time.Now())
	if err != nil {
		logger.Log.Warn("Не удалось добавить пользователя ", zap.Error(err))
		return err
	}
	logger.Log.Info("Добавлен новый пользователь")
	return nil
}

// UserLogin проверяет учетные данные пользователя в базе данных.
func (db *Database) UserLogin(ctx context.Context, login string, password string) error {
	var hashedPassword []byte

	err := db.Conn.QueryRow(ctx, `SELECT password FROM users WHERE login = $1`, login).Scan(&hashedPassword)

	if err == pgx.ErrNoRows {
		return store.ErrAuthentication
	} else if err != nil {
		logger.Log.Warn("Ошибка выполнения запроса ", zap.Error(err))
		return err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return store.ErrAuthentication
	}

	_, err = db.Conn.Exec(ctx, `UPDATE users SET last_time = $1 WHERE login = $2`, time.Now(), login)
	if err != nil {
		logger.Log.Warn("Не удалось обновить значение", zap.Error(err))
		return err
	}

	return nil
}

// AddField добавляет данные.
func (db *Database) AddField(ctx context.Context) bool {
	return true
}

// DelField удаляет данные.
func (db *Database) DelField(ctx context.Context) bool {
	return true
}

// SyncFields синхронизирует данные.
func (db *Database) SyncFields(ctx context.Context, user string, data []*proto.FieldKeep) (remoteData []*proto.FieldKeep, err error) {
	var idUser int
	err = db.Conn.QueryRow(ctx, `SELECT id FROM users WHERE login = $1`, user).Scan(&idUser)
	if err != nil && err != pgx.ErrNoRows {
		logger.Log.Warn("Ошибка выполнения запроса id", zap.Error(err))
		return remoteData, err
	}

	rows, err := db.Conn.Query(ctx, `SELECT uuid, update_at FROM store WHERE user_id = $1`, idUser)
	defer rows.Close()

	checkUpdate := make(map[string]time.Time)
	for rows.Next() {
		var uuid string
		var updateAt time.Time

		err = rows.Scan(&uuid, &updateAt)
		if err != nil {
			logger.Log.Warn("Ошибка при сканировании строки:", zap.Error(err))
			return remoteData, err
		}
		checkUpdate[uuid] = updateAt
	}

	for _, curData := range data {
		val, ok := checkUpdate[curData.GetUuid()]

		if !ok {
			_, err = db.Conn.Exec(ctx,
				`INSERT INTO store (user_id, uuid, update_at) 
				VALUES ($1, $2, $3)`,
				idUser,
				curData.GetUuid(),
				time.Now())
			// _, err = db.Conn.Exec(ctx,
			// 	`INSERT INTO store (user_id, uuid, login, password, data, card_number, card_cvc, card_date, card_owner, update_at) 
			// 	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			// 	idUser,
			// 	curData.GetUuid(),
			// 	curData.GetLogin(),
			// 	curData.GetPassword(),
			// 	curData.GetData(),
			// 	curData.GetCardNumber(),
			// 	curData.GetCardCVC(),
			// 	curData.GetCardDate(),
			// 	curData.GetCardOwner(),
			// 	time.Now())
			if err != nil {
				logger.Log.Warn("Не удалось добавить запись", zap.Error(err))
				return data, err
			}
			logger.Log.Info("Добавлена новая запись")
		} else {
			if val.Sub(time.Unix(curData.GetTimeUpdate(), 0)) > 0 {
				// UPDATE
			}
		}
	}

	return data, err
}
