package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
)

// mainContextTxKey The context used to host the transaction
type mainContextTxKey struct{}

// bizContextTxKey The context used to host the transaction
type bizContextTxKey struct{}

// eventContextTxKey The context used to host the transaction
type eventContextTxKey struct{}

// WithMainTXContext This method creates a new context with the transaction
func WithMainTXContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, mainContextTxKey{}, tx)
}

// GetMainTXByContext This method checks if there is a transaction in the context,
func GetMainTXByContext(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(mainContextTxKey{}).(*gorm.DB)
	return tx, ok
}

// WithBizTXContext This method creates a new context with the transaction
func WithBizTXContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, bizContextTxKey{}, tx)
}

// GetBizTXByContext This method checks if there is a transaction in the context,
func GetBizTXByContext(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(bizContextTxKey{}).(*gorm.DB)
	return tx, ok
}

// WithEventTXContext This method creates a new context with the transaction
func WithEventTXContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, eventContextTxKey{}, tx)
}

// GetEventTXByContext This method checks if there is a transaction in the context,
func GetEventTXByContext(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(eventContextTxKey{}).(*gorm.DB)
	return tx, ok
}

// NewTransaction creates a transaction
func NewTransaction(d *data.Data, logger log.Logger) repository.Transaction {
	return &transactionRepoImpl{
		Data:   d,
		logger: log.NewHelper(log.With(logger, "module", "plugin.gorm")),
	}
}

type transactionRepoImpl struct {
	*data.Data
	logger *log.Helper
}

func (t *transactionRepoImpl) MainExec(ctx context.Context, fn func(ctx context.Context) error) error {
	_, ok := GetMainTXByContext(ctx)
	if ok {
		return fn(ctx)
	}
	return t.GetMainDB().GetDB().Transaction(func(tx *gorm.DB) error {
		txCtx := WithMainTXContext(ctx, tx)
		return fn(txCtx)
	})
}

func (t *transactionRepoImpl) BizExec(ctx context.Context, fn func(ctx context.Context) error) error {
	_, ok := GetBizTXByContext(ctx)
	if ok {
		return fn(ctx)
	}
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return merr.ErrorInternalServerError("team id not found").WithMetadata(map[string]string{"method": "BizExec"})
	}
	bizDB, err := t.GetBizDB(teamID)
	if err != nil {
		return merr.ErrorInternalServerError("biz db not found").WithMetadata(map[string]string{"method": "BizExec"}).WithCause(err)
	}
	return bizDB.GetDB().Transaction(func(tx *gorm.DB) error {
		txCtx := WithBizTXContext(ctx, tx)
		return fn(txCtx)
	})
}

func (t *transactionRepoImpl) EventExec(ctx context.Context, fn func(ctx context.Context) error) error {
	_, ok := GetEventTXByContext(ctx)
	if ok {
		return fn(ctx)
	}
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return merr.ErrorInternalServerError("team id not found").WithMetadata(map[string]string{"method": "EventExec"})
	}
	eventDB, err := t.GetEventDB(teamID)
	if err != nil {
		return merr.ErrorInternalServerError("event db not found").WithMetadata(map[string]string{"method": "EventExec"}).WithCause(err)
	}
	return eventDB.GetDB().Transaction(func(tx *gorm.DB) error {
		txCtx := WithEventTXContext(ctx, tx)
		return fn(txCtx)
	})
}
