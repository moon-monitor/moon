package repository

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type Dashboard interface {
	CreateDashboard(ctx context.Context, dashboard bo.Dashboard) error

	UpdateDashboard(ctx context.Context, dashboard bo.Dashboard) error

	// DeleteDashboard delete dashboard by id
	DeleteDashboard(ctx context.Context, id uint32) error

	// GetDashboard get dashboard by id
	GetDashboard(ctx context.Context, id uint32) (do.Dashboard, error)

	// ListDashboards list dashboards with filter
	ListDashboards(ctx context.Context, req *bo.ListDashboardReq) (*bo.ListDashboardReply, error)

	// BatchUpdateDashboardStatus update multiple dashboards status
	BatchUpdateDashboardStatus(ctx context.Context, req *bo.BatchUpdateDashboardStatusReq) error
}

type DashboardChart interface {
	// SaveDashboardChart save dashboard chart
	// exist id update, else insert
	SaveDashboardChart(ctx context.Context, chart *team.DashboardChart) error

	// DeleteDashboardChart delete dashboard chart by id
	DeleteDashboardChart(ctx context.Context, id uint32) error

	// GetDashboardChart get dashboard chart by id
	GetDashboardChart(ctx context.Context, id uint32) (*team.DashboardChart, error)

	// ListDashboardCharts list dashboard charts with filter
	ListDashboardCharts(ctx context.Context, req *bo.ListDashboardChartReq) (*bo.ListDashboardChartReply, error)

	// BatchUpdateDashboardChartStatus update multiple dashboard charts status
	BatchUpdateDashboardChartStatus(ctx context.Context, ids []uint32, status vobj.GlobalStatus) error
}
