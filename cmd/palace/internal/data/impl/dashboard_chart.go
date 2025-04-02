package impl

import (
	"context"

	"gorm.io/gen/field"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	teamQuery "github.com/moon-monitor/moon/cmd/palace/internal/data/query/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
)

// NewDashboardChartRepo creates a new dashboard chart repository
func NewDashboardChartRepo(data *data.Data, logger log.Logger) repository.DashboardChart {
	return &dashboardChartImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.dashboard_chart")),
	}
}

type dashboardChartImpl struct {
	*data.Data

	helper *log.Helper
}

func (r *dashboardChartImpl) getDashboardChartTX(ctx context.Context) (*teamQuery.Query, error) {
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorInternalServerError("team id not found")
	}
	db, err := r.GetBizDB(teamID)
	if err != nil {
		return nil, err
	}
	return teamQuery.Use(db.GetDB()), nil
}

// SaveDashboardChart save dashboard chart
func (r *dashboardChartImpl) SaveDashboardChart(ctx context.Context, chart *team.DashboardChart) error {
	tx, err := r.getDashboardChartTX(ctx)
	if err != nil {
		return err
	}
	mutation := tx.DashboardChart
	if chart.ID == 0 {
		return mutation.WithContext(ctx).Create(chart)
	}
	updates := []field.AssignExpr{
		mutation.Title.Value(chart.Title),
		mutation.Remark.Value(chart.Remark),
		mutation.Url.Value(chart.Url),
		mutation.Width.Value(chart.Width),
		mutation.Height.Value(chart.Height),
	}
	_, err = mutation.WithContext(ctx).Where(mutation.ID.Eq(chart.ID)).UpdateColumnSimple(updates...)
	return err
}

// DeleteDashboardChart delete dashboard chart by id
func (r *dashboardChartImpl) DeleteDashboardChart(ctx context.Context, id uint32) error {
	tx, err := r.getDashboardChartTX(ctx)
	if err != nil {
		return err
	}
	mutation := tx.DashboardChart
	_, err = mutation.WithContext(ctx).Where(mutation.ID.Eq(id)).Delete()
	return err
}

// GetDashboardChart get dashboard chart by id
func (r *dashboardChartImpl) GetDashboardChart(ctx context.Context, id uint32) (*team.DashboardChart, error) {
	tx, err := r.getDashboardChartTX(ctx)
	if err != nil {
		return nil, err
	}
	mutation := tx.DashboardChart
	return mutation.WithContext(ctx).Where(mutation.ID.Eq(id)).First()
}

// ListDashboardCharts list dashboard charts with filter
func (r *dashboardChartImpl) ListDashboardCharts(ctx context.Context, req *bo.ListDashboardChartReq) (*bo.ListDashboardChartReply, error) {
	tx, err := r.getDashboardChartTX(ctx)
	if err != nil {
		return nil, err
	}
	mutation := tx.DashboardChart
	query := mutation.WithContext(ctx)

	if !req.Status.IsUnknown() {
		query = query.Where(mutation.Status.Eq(req.Status.GetValue()))
	}

	if req.PaginationRequest != nil {
		total, err := query.Count()
		if err != nil {
			return nil, err
		}
		req.WithTotal(total)
		query = query.Offset(req.Offset()).Limit(int(req.Limit))
	}

	charts, err := query.Find()
	if err != nil {
		return nil, err
	}
	return &bo.ListDashboardChartReply{
		PaginationReply: req.ToReply(),
		Charts:          charts,
	}, nil
}

// BatchUpdateDashboardChartStatus update multiple dashboard charts status
func (r *dashboardChartImpl) BatchUpdateDashboardChartStatus(ctx context.Context, ids []uint32, status vobj.GlobalStatus) error {
	tx, err := r.getDashboardChartTX(ctx)
	if err != nil {
		return err
	}
	mutation := tx.DashboardChart
	_, err = mutation.WithContext(ctx).Where(mutation.ID.In(ids...)).UpdateColumnSimple(mutation.Status.Value(int8(status)))
	return err
}
