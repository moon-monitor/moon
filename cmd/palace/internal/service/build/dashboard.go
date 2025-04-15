package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

// ToDashboardItemProto converts a business object to a proto object
func ToDashboardItemProto(dashboard do.Dashboard) *common.TeamDashboardItem {
	if dashboard == nil {
		return nil
	}

	return &common.TeamDashboardItem{
		Id:        dashboard.GetDashboardID(),
		Name:      dashboard.GetTitle(),
		Remark:    dashboard.GetRemark(),
		Status:    common.GlobalStatus(dashboard.GetStatus()),
		ColorHex:  dashboard.GetColorHex(),
		CreatedAt: dashboard.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: dashboard.GetUpdatedAt().Format(time.DateTime),
	}
}

// ToDashboardItemProtoList converts a list of business objects to proto objects
func ToDashboardItemProtoList(dashboards []*team.Dashboard) []*common.TeamDashboardItem {
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
func ToDashboardChartItemProto(chart *team.DashboardChart) *common.TeamDashboardChartItem {
	if chart == nil {
		return nil
	}

	return &common.TeamDashboardChartItem{
		Id:          chart.ID,
		DashboardID: chart.DashboardID,
		Title:       chart.Title,
		Remark:      chart.Remark,
		Status:      common.GlobalStatus(chart.Status),
		Url:         chart.Url,
		Width:       chart.Width,
		Height:      chart.Height,
		CreatedAt:   chart.CreatedAt.Format(time.DateTime),
		UpdatedAt:   chart.UpdatedAt.Format(time.DateTime),
	}
}

// ToDashboardChartItemProtoList converts a list of business objects to proto objects
func ToDashboardChartItemProtoList(charts []*team.DashboardChart) []*common.TeamDashboardChartItem {
	if charts == nil {
		return nil
	}

	items := make([]*common.TeamDashboardChartItem, 0, len(charts))
	for _, chart := range charts {
		items = append(items, ToDashboardChartItemProto(chart))
	}

	return items
}
