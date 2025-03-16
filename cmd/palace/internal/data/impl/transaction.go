package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"gorm.io/gorm"

	gormdb "github.com/moon-monitor/moon/pkg/plugin/gorm"
)

// NewTransaction creates a transaction
func NewTransaction(d *data.Data, logger log.Logger) repository.Transaction {
	return &transactionRepoImpl{
		tx:     d.GetMainDB(),
		logger: log.NewHelper(log.With(logger, "module", "plugin.gorm")),
	}
}

type transactionRepoImpl struct {
	tx     gormdb.DB
	logger *log.Helper
}

// contextTxKey The context used to host the transaction
type contextTxKey struct{}

func (t *transactionRepoImpl) MainExec(ctx context.Context, fn func(ctx context.Context) error) error {
	_, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return fn(ctx)
	}
	return t.tx.GetTx(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, contextTxKey{}, tx)
		return fn(txCtx)
	})
}

func (t *transactionRepoImpl) BizExec(ctx context.Context, fn func(ctx context.Context) error) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionRepoImpl) EventExec(ctx context.Context, fn func(ctx context.Context) error) error {
	//TODO implement me
	panic("implement me")
}
