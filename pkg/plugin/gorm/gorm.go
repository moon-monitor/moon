package gorm

import (
	"context"

	"gorm.io/gorm"

	"github.com/moon-monitor/moon/pkg/config"
)

type DB interface {
	GetTx(ctx context.Context) *gorm.DB
	Close() error
}

// NewDB creates a new DB instance
func NewDB(c *config.Database) (DB, error) {
	// TODO create a new DB instance
	return &db{}, nil
}

type db struct {
	*gorm.DB
}

// GetTx This method checks if there is a transaction in the context,
// and if so returns the client with the transaction
func (t *db) GetTx(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return t.DB
}

// Close This method closes the DB instance
func (t *db) Close() error {
	s, err := t.DB.DB()
	if err != nil {
		return err
	}
	return s.Close()
}
