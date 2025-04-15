package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
)

// DashboardBiz is a dashboard business logic implementation.
type DashboardBiz struct {
	dashboardRepo      repository.Dashboard
	dashboardChartRepo repository.DashboardChart
	log                *log.Helper
}

// NewDashboardBiz creates a new dashboard business logic.
func NewDashboardBiz(
	dashboardRepo repository.Dashboard,
	dashboardChartRepo repository.DashboardChart,
	logger log.Logger,
) *DashboardBiz {
	return &DashboardBiz{
		dashboardRepo:      dashboardRepo,
		dashboardChartRepo: dashboardChartRepo,
		log:                log.NewHelper(log.With(logger, "module", "biz.dashboard")),
	}
}

// SaveDashboard saves a dashboard.
func (b *DashboardBiz) SaveDashboard(ctx context.Context, req *bo.SaveDashboardReq) error {
	if req.GetID() == 0 {
		return b.dashboardRepo.CreateDashboard(ctx, req)
	}

	dashboardDo, err := b.dashboardRepo.GetDashboard(ctx, req.GetID())
	if err != nil {
		return err
	}

	dashboard := req.WithDashboard(dashboardDo)

	return b.dashboardRepo.UpdateDashboard(ctx, dashboard)
}

// DeleteDashboard deletes a dashboard.
func (b *DashboardBiz) DeleteDashboard(ctx context.Context, id uint32) error {
	return b.dashboardRepo.DeleteDashboard(ctx, id)
}

// GetDashboard gets a dashboard.
func (b *DashboardBiz) GetDashboard(ctx context.Context, id uint32) (do.Dashboard, error) {
	dashboard, err := b.dashboardRepo.GetDashboard(ctx, id)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}

// ListDashboard lists dashboards.
func (b *DashboardBiz) ListDashboard(ctx context.Context, req *bo.ListDashboardReq) (*bo.ListDashboardReply, error) {
	reply, err := b.dashboardRepo.ListDashboards(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// BatchUpdateDashboardStatus updates multiple dashboards' status.
func (b *DashboardBiz) BatchUpdateDashboardStatus(ctx context.Context, req *bo.BatchUpdateDashboardStatusReq) error {
	return b.dashboardRepo.BatchUpdateDashboardStatus(ctx, req)
}

// SaveDashboardChart saves a dashboard chart.
func (b *DashboardBiz) SaveDashboardChart(ctx context.Context, req *bo.SaveDashboardChartReq) error {
	if req.GetID() == 0 {
		return b.dashboardChartRepo.CreateDashboardChart(ctx, req)
	}

	dashboardChartDo, err := b.dashboardChartRepo.GetDashboardChart(ctx, req.GetID())
	if err != nil {
		return err
	}

	chart := req.WithDashboardChart(dashboardChartDo)

	return b.dashboardChartRepo.UpdateDashboardChart(ctx, chart)
}

// DeleteDashboardChart deletes a dashboard chart.
func (b *DashboardBiz) DeleteDashboardChart(ctx context.Context, id uint32) error {
	return b.dashboardChartRepo.DeleteDashboardChart(ctx, id)
}

// GetDashboardChart gets a dashboard chart.
func (b *DashboardBiz) GetDashboardChart(ctx context.Context, id uint32) (do.DashboardChart, error) {
	chart, err := b.dashboardChartRepo.GetDashboardChart(ctx, id)
	if err != nil {
		return nil, err
	}

	return chart, nil
}

// ListDashboardChart lists dashboard charts.
func (b *DashboardBiz) ListDashboardChart(ctx context.Context, req *bo.ListDashboardChartReq) (*bo.ListDashboardChartReply, error) {
	reply, err := b.dashboardChartRepo.ListDashboardCharts(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// BatchUpdateDashboardChartStatus updates multiple dashboard charts' status.
func (b *DashboardBiz) BatchUpdateDashboardChartStatus(ctx context.Context, req *bo.BatchUpdateDashboardChartStatusReq) error {
	return b.dashboardChartRepo.BatchUpdateDashboardChartStatus(ctx, req)
}
