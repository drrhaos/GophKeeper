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
	AddField(ctx context.Context, user string, data *proto.FieldKeep) (string, bool)
	EditField(ctx context.Context, user string, uuid string, data *proto.FieldKeep) (*proto.FieldKeep, bool)
	DelField(ctx context.Context, user string, uuid string) (string, bool)
	ListFields(ctx context.Context, login string) ([]*proto.FieldExtended, bool)
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
func (sc *StorageContext) UserRegister(ctx context.Context, user string, password string) error {
	return sc.storage.UserRegister(ctx, user, password)
}

// UserLogin проверяет учетные данные пользователя.
func (sc *StorageContext) UserLogin(ctx context.Context, user string, password string) error {
	return sc.storage.UserLogin(ctx, user, password)
}

// AddField добавляет данные.
func (sc *StorageContext) AddField(ctx context.Context, user string, data *proto.FieldKeep) (string, bool) {
	return sc.storage.AddField(ctx, user, data)
}

// EditField изменяте данные.
func (sc *StorageContext) EditField(ctx context.Context, user string, uuid string, data *proto.FieldKeep) (*proto.FieldKeep, bool) {
	return sc.storage.EditField(ctx, user, uuid, data)
}

// DelField удаляет данные.
func (sc *StorageContext) DelField(ctx context.Context, user string, uuid string) (string, bool) {
	return sc.storage.DelField(ctx, user, uuid)
}

// ListFields возвращает список данных пользователя.
func (sc *StorageContext) ListFields(ctx context.Context, user string) ([]*proto.FieldExtended, bool) {
	return sc.storage.ListFields(ctx, user)
}
