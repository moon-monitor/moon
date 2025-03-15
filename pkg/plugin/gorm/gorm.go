package gorm

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/merr"
)

type DB interface {
	GetTx(ctx context.Context) *gorm.DB
	Close() error
}

// NewDB creates a new DB instance
func NewDB(c *config.Database) (DB, error) {
	var opts []gorm.Option
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	opts = append(opts, gormConfig)
	var dialector gorm.Dialector
	dsn := c.GetDsn()
	drive := c.GetDriver()
	switch drive {
	case config.Database_MYSQL:
		dialector = mysql.Open(dsn)
	case config.Database_SQLITE:
		dialector = sqlite.Open(dsn)
	default:
		return nil, merr.ErrorInternalServerError("invalid driver: %s", drive)
	}
	conn, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, merr.ErrorInternalServerError("connect db error: %s", err)
	}
	if drive == config.Database_SQLITE {
		_ = conn.Exec("PRAGMA journal_mode=WAL;")
	}
	if c.GetDebug() {
		conn = conn.Debug()
	}
	return &db{DB: conn}, nil
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
