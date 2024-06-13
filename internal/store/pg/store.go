// Package pg реализует взаимодействие с базой данной Postgres.
package pg

import (
	"context"
	"errors"
	"gophkeeper/internal/logger"
	"gophkeeper/pkg/proto"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
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


// AddField добавляет данные.
func (sc *Database) AddField(ctx context.Context) bool {
	return true
}

// DelField удаляет данные.
func (sc *Database) DelField(ctx context.Context) bool {
	return true
}

// SyncFields синхронизирует данные.
func (sc *Database) SyncFields(ctx context.Context, data []proto.FieldKeep) ([]proto.FieldKeep, error) {

	return data, nil
}