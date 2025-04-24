package impl

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewTeamStrategyRepo(data *data.Data) repository.TeamStrategy {
	return &teamStrategyImpl{
		Data: data,
	}
}

type teamStrategyImpl struct {
	*data.Data
}

// List implements repository.TeamStrategy.
func (t *teamStrategyImpl) List(ctx context.Context, params *bo.ListTeamStrategyParams) (*bo.ListTeamStrategyReply, error) {
	tx, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}
	query := tx.Strategy
	wrappers := query.WithContext(ctx).Where(query.TeamID.Eq(teamId))
	if validate.TextIsNotNull(params.Keyword) {
		wrappers = wrappers.Where(query.Name.Like("%" + params.Keyword + "%"))
	}
	if len(params.Status) > 0 {
		wrappers = wrappers.Where(query.Status.In(slices.Map(params.Status, func(status vobj.GlobalStatus) int8 { return status.GetValue() })...))
	}
	if validate.IsNotNil(params.PaginationRequest) {
		total, err := wrappers.Count()
		if err != nil {
			return nil, err
		}
		params.WithTotal(total)
		wrappers = wrappers.Limit(int(params.Limit)).Offset(params.Offset())
	}
	strategies, err := wrappers.Find()
	if err != nil {
		return nil, err
	}
	return params.ToListTeamStrategyReply(strategies), nil
}

// Subscribe implements repository.TeamStrategy.
func (t *teamStrategyImpl) Subscribe(ctx context.Context, params *bo.ToSubscribeTeamStrategyParams) error {
	tx, _, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	subscriberDo := &team.StrategySubscriber{
		StrategyID:    params.StrategyID,
		SubscribeType: params.SubscribeType,
	}

	subscriberDo.WithContext(ctx)

	return tx.StrategySubscriber.WithContext(ctx).Create(subscriberDo)
}

// SubscribeList implements repository.TeamStrategy.
func (t *teamStrategyImpl) SubscribeList(ctx context.Context, params *bo.ToSubscribeTeamStrategiesParams) (*bo.ToSubscribeTeamStrategiesReply, error) {
	tx, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}
	query := tx.StrategySubscriber
	wrappers := query.WithContext(ctx).Where(query.TeamID.Eq(teamId))
	if params.StrategyID > 0 {
		wrappers = wrappers.Where(query.StrategyID.Eq(params.StrategyID))
	}
	if len(params.Subscribers) > 0 {
		wrappers = wrappers.Where(query.CreatorID.In(params.Subscribers...))
	}
	if validate.IsNotNil(params.PaginationRequest) {
		total, err := wrappers.Count()
		if err != nil {
			return nil, err
		}
		params.WithTotal(total)
		wrappers = wrappers.Limit(int(params.Limit)).Offset(params.Offset())
	}
	subscribers, err := wrappers.Find()
	if err != nil {
		return nil, err
	}

	return params.ToSubscribeTeamStrategiesReply(subscribers), nil
}

// UpdateStatus implements repository.TeamStrategy.
func (t *teamStrategyImpl) UpdateStatus(ctx context.Context, params *bo.UpdateTeamStrategiesStatusParams) error {
	tx, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	mutation := tx.Strategy
	wrappers := []gen.Condition{
		mutation.ID.In(params.StrategyIds...),
		mutation.TeamID.Eq(teamId),
	}
	mutations := []field.AssignExpr{
		mutation.Status.Value(params.Status.GetValue()),
	}
	_, err = mutation.WithContext(ctx).Where(wrappers...).UpdateSimple(mutations...)
	return err
}
