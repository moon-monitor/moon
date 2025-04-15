package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

// ToDashboardItemProto converts a business object to a proto object
func ToDashboardItemProto(dashboard do.Dashboard) *common.TeamDashboardItem {
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

// ToDashboardItemProtoList converts a list of business objects to proto objects
func ToDashboardItemProtoList(dashboards []do.Dashboard) []*common.TeamDashboardItem {
	if dashboards == nil {
		return nil
	}

	items := make([]*common.TeamDashboardItem, 0, len(dashboards))
	for _, dashboard := range dashboards {
		items = append(items, ToDashboardItemProto(dashboard))
	}

	return items
}

// ToDashboardChartItemProto converts a business object to a proto object
func ToDashboardChartItemProto(chart do.DashboardChart) *common.TeamDashboardChartItem {
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

// ToDashboardChartItemProtoList converts a list of business objects to proto objects
func ToDashboardChartItemProtoList(charts []do.DashboardChart) []*common.TeamDashboardChartItem {
	if charts == nil {
		return nil
	}

	items := make([]*common.TeamDashboardChartItem, 0, len(charts))
	for _, chart := range charts {
		items = append(items, ToDashboardChartItemProto(chart))
	}

	return items
}
