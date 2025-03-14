package gorm

import (
	"context"

	"gorm.io/gorm"
)

func NewDB() (*DB, error) {
	return &DB{}, nil
}

type DB struct {
	*gorm.DB
}

// Config GORM数据库配置
type Config interface {
	GetDriver() string
	GetDsn() string
	GetDebug() bool
}

// GetTx This method checks if there is a transaction in the context, and if so returns the client with the transaction
func (t *DB) GetTx(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return t.DB
}
