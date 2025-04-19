package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

// ToDashboardItem converts a business object to a proto object
func ToDashboardItem(dashboard do.Dashboard) *common.TeamDashboardItem {
	if dashboard == nil {
		return nil
	}

	return &common.TeamDashboardItem{
		Id:        dashboard.GetID(),
		Name:      dashboard.GetTitle(),
		Remark:    dashboard.GetRemark(),
		Status:    common.GlobalStatus(dashboard.GetStatus()),
		ColorHex:  dashboard.GetColorHex(),
		CreatedAt: dashboard.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: dashboard.GetUpdatedAt().Format(time.DateTime),
	}
}

// ToDashboardItems converts multiple business objects to proto objects
func ToDashboardItems(dashboards []do.Dashboard) []*common.TeamDashboardItem {
	return slices.Map(dashboards, ToDashboardItem)
}

// ToDashboardChartItem converts a business object to a proto object
func ToDashboardChartItem(chart do.DashboardChart) *common.TeamDashboardChartItem {
	if chart == nil {
		return nil
	}

	return &common.TeamDashboardChartItem{
		Id:          chart.GetID(),
		DashboardID: chart.GetDashboardID(),
		Title:       chart.GetTitle(),
		Remark:      chart.GetRemark(),
		Status:      common.GlobalStatus(chart.GetStatus()),
		Url:         chart.GetUrl(),
		Width:       chart.GetWidth(),
		Height:      chart.GetHeight(),
		CreatedAt:   chart.GetCreatedAt().Format(time.DateTime),
		UpdatedAt:   chart.GetUpdatedAt().Format(time.DateTime),
	}
}

// ToDashboardChartItems converts multiple business objects to proto objects
func ToDashboardChartItems(charts []do.DashboardChart) []*common.TeamDashboardChartItem {
	return slices.Map(charts, ToDashboardChartItem)
}
