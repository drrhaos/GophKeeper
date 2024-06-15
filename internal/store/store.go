// Package store предназначен для хранения метрик.
package store

import (
	"context"
	"errors"

	"gophkeeper/pkg/proto"
)

var (
	ErrLoginDuplicate = errors.New("user duplicate")                // ErrLoginDuplicate пользователь уже существует.
	ErrAuthentication = errors.New("invalid user name or password") // ErrAuthentication ошибка Authentication.
)

// StorageInterface описывает набор методов которые должны реализовывать хранилища.
type StorageInterface interface {
	UserRegister(ctx context.Context, login string, password string) error
	UserLogin(ctx context.Context, login string, password string) error
	AddField(ctx context.Context) bool
	DelField(ctx context.Context) bool
	SyncFields(ctx context.Context, user string, data []*proto.FieldKeep) ([]*proto.FieldKeep, error)
}

// StorageContext содержит текущее хранилище.
type StorageContext struct {
	storage StorageInterface
}

// SetStorage устанавливает хранилище.
func (sc *StorageContext) SetStorage(storage StorageInterface) {
	sc.storage = storage
}

// UserRegister добавляет нового пользоватяля.
func (sc *StorageContext) UserRegister(ctx context.Context, login string, password string) error {
	return sc.storage.UserRegister(ctx, login, password)
}

// UserLogin проверяет учетные данные пользователя.
func (sc *StorageContext) UserLogin(ctx context.Context, login string, password string) error {
	return sc.storage.UserLogin(ctx, login, password)
}

// AddField добавляет данные.
func (sc *StorageContext) AddField(ctx context.Context) bool {
	return sc.storage.AddField(ctx)
}

// DelField удаляет данные.
func (sc *StorageContext) DelField(ctx context.Context) bool {
	return sc.storage.DelField(ctx)
}

// SyncFields синхронизирует данные.
func (sc *StorageContext) SyncFields(ctx context.Context, user string, data []*proto.FieldKeep) ([]*proto.FieldKeep, error) {
	return sc.storage.SyncFields(ctx, user, data)
}
