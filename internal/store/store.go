// Package store предназначен для хранения метрик.
package store

import (
	"context"
	"gophkeeper/pkg/proto"
)

// StorageInterface описывает набор методов которые должны реализовывать хранилища.
type StorageInterface interface {
	AddField(ctx context.Context) bool
	DelField(ctx context.Context) bool
	SyncFields(ctx context.Context, data []proto.FieldKeep) ([]proto.FieldKeep, error)
}

// StorageContext содержит текущее хранилище.
type StorageContext struct {
	storage StorageInterface
}

// SetStorage устанавливает хранилище.
func (sc *StorageContext) SetStorage(storage StorageInterface) {
	sc.storage = storage
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
func (sc *StorageContext) SyncFields(ctx context.Context, data []proto.FieldKeep) ([]proto.FieldKeep, error) {
	return sc.storage.SyncFields(ctx, data)
}