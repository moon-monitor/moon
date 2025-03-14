package gorm

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Transaction interface {
	Exec(ctx context.Context, fn func(ctx context.Context) error) error
}

type transactionRepoImpl struct {
	tx     DB
	logger *log.Helper
}

// contextTxKey The context used to host the transaction
type contextTxKey struct{}

// NewTransaction creates a transaction
func NewTransaction(db DB, logger log.Logger) Transaction {
	return &transactionRepoImpl{
		tx:     db,
		logger: log.NewHelper(log.With(logger, "module", "plugin.gorm")),
	}
}

func (t *transactionRepoImpl) Exec(ctx context.Context, fn func(ctx context.Context) error) error {
	_, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return fn(ctx)
	}
	return t.tx.GetTx(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, contextTxKey{}, tx)
		return fn(txCtx)
	})
}
