package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/util/crypto"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewTeamMetricDatasourceRepo(data *data.Data, logger log.Logger) repository.TeamDatasourceMetric {
	return &teamMetricDatasourceImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.team.datasource.metric")),
	}
}

type teamMetricDatasourceImpl struct {
	*data.Data
	helper *log.Helper
}

func (t *teamMetricDatasourceImpl) Create(ctx context.Context, req *bo.SaveTeamMetricDatasource) error {
	metricDatasourceDo := &team.DatasourceMetric{
		Name:           req.Name,
		Status:         req.Status,
		Remark:         req.Remark,
		Driver:         req.Driver,
		Endpoint:       req.Endpoint,
		ScrapeInterval: req.ScrapeInterval,
		Headers:        req.Headers,
		QueryMethod:    req.QueryMethod,
		CA:             req.CA,
		TLS:            crypto.NewObject(req.TLS),
		BasicAuth:      crypto.NewObject(req.BasicAuth),
		Extra:          req.Extra,
	}
	metricDatasourceDo.WithContext(ctx)
	bizMutation, _, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	return bizMutation.DatasourceMetric.WithContext(ctx).Create(metricDatasourceDo)
}

func (t *teamMetricDatasourceImpl) Update(ctx context.Context, req *bo.SaveTeamMetricDatasource) error {
	bizMutation, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	mutation := bizMutation.DatasourceMetric
	wrapper := []gen.Condition{
		mutation.TeamID.Eq(teamId),
		mutation.ID.Eq(req.ID),
	}
	mutations := []field.AssignExpr{
		mutation.Name.Value(req.Name),
		mutation.Status.Value(req.Status.GetValue()),
		mutation.Remark.Value(req.Remark),
		mutation.Driver.Value(req.Driver.GetValue()),
		mutation.Endpoint.Value(req.Endpoint),
		mutation.ScrapeInterval.Value(int64(req.ScrapeInterval)),
		mutation.Headers.Value(req.Headers),
		mutation.QueryMethod.Value(req.QueryMethod.GetValue()),
		mutation.CA.Value(req.CA),
		mutation.TLS.Value(crypto.NewObject(req.TLS)),
		mutation.BasicAuth.Value(crypto.NewObject(req.BasicAuth)),
		mutation.Extra.Value(req.Extra),
	}
	_, err = mutation.WithContext(ctx).Where(wrapper...).UpdateSimple(mutations...)
	return err

}

func (t *teamMetricDatasourceImpl) UpdateStatus(ctx context.Context, req *bo.UpdateTeamMetricDatasourceStatusRequest) error {
	bizMutation, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	mutation := bizMutation.DatasourceMetric
	wrapper := []gen.Condition{
		mutation.TeamID.Eq(teamId),
		mutation.ID.Eq(req.DatasourceID),
	}
	_, err = mutation.WithContext(ctx).Where(wrapper...).UpdateSimple(mutation.Status.Value(req.Status.GetValue()))
	return err
}

func (t *teamMetricDatasourceImpl) Delete(ctx context.Context, datasourceID uint32) error {
	bizMutation, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	mutation := bizMutation.DatasourceMetric
	wrapper := []gen.Condition{
		mutation.TeamID.Eq(teamId),
		mutation.ID.Eq(datasourceID),
	}
	_, err = mutation.WithContext(ctx).Where(wrapper...).Delete()
	return err
}

func (t *teamMetricDatasourceImpl) Get(ctx context.Context, datasourceID uint32) (do.DatasourceMetric, error) {
	bizQuery, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}
	mutation := bizQuery.DatasourceMetric
	wrapper := []gen.Condition{
		mutation.TeamID.Eq(teamId),
		mutation.ID.Eq(datasourceID),
	}
	return mutation.WithContext(ctx).Where(wrapper...).First()
}

func (t *teamMetricDatasourceImpl) List(ctx context.Context, req *bo.ListTeamMetricDatasource) (*bo.ListTeamMetricDatasourceReply, error) {
	bizQuery, teamId, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}
	mutation := bizQuery.DatasourceMetric
	wrapper := mutation.WithContext(ctx).Where(mutation.TeamID.Eq(teamId))

	if !req.Status.IsUnknown() {
		wrapper = wrapper.Where(mutation.Status.Eq(req.Status.GetValue()))
	}
	if !validate.TextIsNull(req.Keyword) {
		ors := []gen.Condition{
			mutation.Name.Like(req.Keyword),
			mutation.Remark.Like(req.Keyword),
			mutation.Endpoint.Like(req.Keyword),
		}
		wrapper = wrapper.Where(mutation.Or(ors...))
	}
	if validate.IsNotNil(req.PaginationRequest) {
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		wrapper = wrapper.Offset(req.Offset()).Limit(int(req.Limit))
		req.WithTotal(total)
	}
	datasourceDos, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	return req.ToListTeamMetricDatasourceReply(datasourceDos), nil
}
