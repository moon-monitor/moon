package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/teamgen"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
	"gorm.io/gen/field"
)

func NewTeamDictRepo(d *data.Data, logger log.Logger) repository.TeamDict {
	return &teamDictImpl{
		Data:   d,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.team_dict")),
	}
}

type teamDictImpl struct {
	*data.Data
	helper *log.Helper
}

func (t *teamDictImpl) bizQuery(teamID uint32) (*teamgen.Query, error) {
	bizDB, err := t.GetBizDB(teamID)
	if err != nil {
		return nil, err
	}
	return teamgen.Use(bizDB.GetDB()), nil
}

func (t *teamDictImpl) Get(ctx context.Context, teamID, dictID uint32) (bo.Dict, error) {
	bizQuery, err := t.bizQuery(teamID)
	if err != nil {
		return nil, err
	}
	bizDictQuery := bizQuery.Dict
	return bizDictQuery.WithContext(ctx).Where(bizDictQuery.TeamID.Eq(teamID)).Where(bizDictQuery.ID.Eq(dictID)).First()
}

func (t *teamDictImpl) Delete(ctx context.Context, teamID, dictID uint32) error {
	bizQuery, err := t.bizQuery(teamID)
	if err != nil {
		return err
	}
	bizDictQuery := bizQuery.Dict
	_, err = bizDictQuery.WithContext(ctx).Where(bizDictQuery.ID.Eq(dictID)).Delete()
	return err
}

func (t *teamDictImpl) Create(ctx context.Context, teamID uint32, dict bo.Dict) error {
	bizQuery, err := t.bizQuery(teamID)
	if err != nil {
		return err
	}
	dictDo := &team.Dict{
		Key:      dict.GetKey(),
		Value:    dict.GetValue(),
		Lang:     dict.GetLang(),
		Color:    dict.GetColor(),
		DictType: dict.GetType(),
		Status:   dict.GetStatus(),
	}
	dictDo.WithContext(ctx)
	bizDictQuery := bizQuery.Dict
	return bizDictQuery.WithContext(ctx).Create(dictDo)
}

func (t *teamDictImpl) Update(ctx context.Context, teamID uint32, dict bo.Dict) error {
	bizQuery, err := t.bizQuery(teamID)
	if err != nil {
		return err
	}
	bizDictQuery := bizQuery.Dict
	mutations := []field.AssignExpr{
		bizDictQuery.Key.Value(dict.GetKey()),
		bizDictQuery.Value.Value(dict.GetValue()),
		bizDictQuery.Lang.Value(dict.GetLang()),
		bizDictQuery.Color.Value(dict.GetColor()),
		bizDictQuery.DictType.Value(dict.GetType().GetValue()),
		bizDictQuery.Status.Value(dict.GetStatus().GetValue()),
	}
	_, err = bizDictQuery.WithContext(ctx).
		Where(bizDictQuery.TeamID.Eq(teamID)).
		Where(bizDictQuery.ID.Eq(dict.GetID())).
		UpdateColumnSimple(mutations...)
	return err
}

func (t *teamDictImpl) UpdateStatus(ctx context.Context, teamID uint32, dictIds []uint32, status vobj.GlobalStatus) error {
	bizQuery, err := t.bizQuery(teamID)
	if err != nil {
		return err
	}
	bizDictQuery := bizQuery.Dict
	_, err = bizDictQuery.WithContext(ctx).
		Where(bizDictQuery.TeamID.Eq(teamID)).
		Where(bizDictQuery.ID.In(dictIds...)).
		UpdateColumnSimple(bizDictQuery.Status.Value(status.GetValue()))
	return err
}

func (t *teamDictImpl) List(ctx context.Context, teamID uint32, req *bo.ListDictReq) (*bo.ListDictReply, error) {
	bizQuery, err := t.bizQuery(teamID)
	if err != nil {
		return nil, err
	}
	bizDictQuery := bizQuery.Dict
	wrapper := bizDictQuery.WithContext(ctx)
	if len(req.Langs) > 0 {
		wrapper = wrapper.Where(bizDictQuery.Lang.In(req.Langs...))
	}
	if len(req.DictTypes) > 0 {
		dictTypes := slices.Map(req.DictTypes, func(item vobj.DictType) int8 { return item.GetValue() })
		wrapper = wrapper.Where(bizDictQuery.DictType.In(dictTypes...))
	}
	if !req.Status.IsUnknown() {
		wrapper = wrapper.Where(bizDictQuery.Status.Eq(req.Status.GetValue()))
	}
	if !validate.TextIsNull(req.Keyword) {
		wrapper = wrapper.Where(bizDictQuery.Key.Like(req.Keyword))
	}
	if validate.IsNotNil(req.PaginationRequest) {
		wrapper = wrapper.Offset(req.Offset()).Limit(int(req.Limit))
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		req.WithTotal(total)
	}
	dictItems, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	return req.NewListDictReply(slices.Map(dictItems, func(item *team.Dict) bo.Dict { return item })), nil
}
