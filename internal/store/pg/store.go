// Package pg реализует взаимодействие с базой данной Postgres.
package pg

import (
	"context"
	"embed"
	"errors"
	"time"

	"gophkeeper/internal/logger"
	"gophkeeper/internal/server/configure"
	"gophkeeper/internal/store"
	"gophkeeper/pkg/proto"

	"github.com/golang-migrate/migrate/v4"
	// необходима для миграции
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

//go:embed migrations
var fs embed.FS

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
func Migrations(cfg configure.Config) error {
	source, _ := iofs.New(fs, "migrations")

	m, err := migrate.NewWithSourceInstance("iofs", source, cfg.DatabaseDsn)
	if err != nil {
		logger.Log.Warn("Не удалось подключиться к базе данных", zap.Error(err))
		return err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Log.Warn("Не удалось выполнить миграцию", zap.Error(err))
		return err
	}
	logger.Log.Info("Миграции успешно применены")
	return nil
}

func (db *Database) getUserID(ctx context.Context, user string) (idUser int, err error) {
	err = db.Conn.QueryRow(ctx, `SELECT id FROM users WHERE login = $1`, user).Scan(&idUser)
	if err != nil && err != pgx.ErrNoRows {
		logger.Log.Warn("Ошибка выполнения запроса id", zap.Error(err))
		return idUser, err
	}

	return idUser, err
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
func (db *Database) AddField(ctx context.Context, user string, data *proto.FieldKeep) (string, *proto.FieldKeep, bool) {
	uuid := uuid.New().String()

	idUser, err := db.getUserID(ctx, user)
	if err != nil {
		return "", nil, false
	}
	_, err = db.Conn.Exec(ctx,
		`INSERT INTO store (user_id, uuid, login, password, data, card_number, card_cvc, card_date, card_owner, update_at, file_name)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		idUser,
		uuid,
		data.GetLogin(),
		data.GetPassword(),
		data.GetData(),
		data.GetCardNumber(),
		data.GetCardCVC(),
		data.GetCardDate(),
		data.GetCardOwner(),
		time.Now(),
		data.GetFileName())
	if err != nil {
		logger.Log.Warn("Не удалось добавить запись", zap.Error(err))
		return "", nil, false
	}
	logger.Log.Info("Добавлена новая запись")

	return uuid, data, true
}

// EditField добавляет данные.
func (db *Database) EditField(ctx context.Context, user string, uuid string, data *proto.FieldKeep) (*proto.FieldKeep, bool) {
	idUser, err := db.getUserID(ctx, user)
	if err != nil {
		return nil, false
	}
	_, err = db.Conn.Exec(ctx,
		`UPDATE store 
			SET
				login=$1,
				password=$2, 
				data=$3,
				card_number=$4,
				card_cvc=$5,
				card_date=$6,
				card_owner=$7,
				file_name=$8,
				update_at=$9
			WHERE 
				user_id=$10
			AND
				uuid=$11`,
		data.GetLogin(),
		data.GetPassword(),
		data.GetData(),
		data.GetCardNumber(),
		data.GetCardCVC(),
		data.GetCardDate(),
		data.GetCardOwner(),
		data.GetFileName(),
		time.Now(),
		idUser,
		uuid)
	if err != nil {
		logger.Log.Warn("Не удалось изменить запись", zap.Error(err))
		return nil, false
	}
	logger.Log.Info("Запись изменена")

	return data, true
}

// DelField удаляет данные.
func (db *Database) DelField(ctx context.Context, user string, uuid string) (string, bool) {
	idUser, err := db.getUserID(ctx, user)
	if err != nil {
		return "", false
	}
	_, err = db.Conn.Exec(ctx,
		`DELETE FROM store WHERE user_id=$1 AND uuid=$2`,
		idUser,
		uuid)
	if err != nil {
		logger.Log.Warn("Не удалось удалить запись", zap.Error(err))
		return "", false
	}
	logger.Log.Info("Запись удалена")

	return uuid, true
}

// ListFields возвращает список данных.
func (db *Database) ListFields(ctx context.Context, user string) (fieldExt *proto.ListFielsdKeepResponse, ok bool) {
	fieldExt = &proto.ListFielsdKeepResponse{
		Data: make(map[string]*proto.FieldKeep),
	}

	idUser, err := db.getUserID(ctx, user)
	if err != nil {
		return fieldExt, false
	}

	rows, err := db.Conn.Query(ctx,
		`SELECT 
			uuid,
			login,
			password,
			data,
			card_number,
			card_cvc,
			card_date,
			card_owner,
			file_name
		FROM 
			store
		WHERE 
			user_id = $1`,
		idUser)
	if err != nil {
		logger.Log.Warn("Ошибка выполнения запроса ", zap.Error(err))
		return fieldExt, false
	}

	defer rows.Close()

	for rows.Next() {
		var field proto.FieldKeep
		// err = rows.Scan(&field.Uuid)
		var uuid string
		err = rows.Scan(
			&uuid,
			&field.Login,
			&field.Password,
			&field.Data,
			&field.CardNumber,
			&field.CardCVC,
			&field.CardDate,
			&field.CardOwner,
			&field.FileName)
		if err != nil {
			logger.Log.Warn("Ошибка при сканировании строки:", zap.Error(err))
			return fieldExt, false
		}

		fieldExt.Data[uuid] = &field
	}

	return fieldExt, true
}
