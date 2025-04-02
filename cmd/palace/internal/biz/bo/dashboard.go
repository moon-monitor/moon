package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

// SaveDashboardReq represents a request to save a dashboard
type SaveDashboardReq struct {
	ID       uint32
	Title    string
	Remark   string
	Status   vobj.GlobalStatus
	ColorHex string
}

// SaveDashboardChartReq represents a request to save a dashboard chart
type SaveDashboardChartReq struct {
	ID          uint32
	DashboardID uint32
	Title       string
	Remark      string
	Status      vobj.GlobalStatus
	Url         string
	Width       string
	Height      string
}

// ListDashboardReq represents a request to list dashboards
type ListDashboardReq struct {
	*PaginationRequest
	Status vobj.GlobalStatus
}

// ListDashboardReply represents a reply to list dashboards
type ListDashboardReply struct {
	*PaginationReply
	Dashboards []*team.Dashboard
}

// ListDashboardChartReq represents a request to list dashboard charts
type ListDashboardChartReq struct {
	*PaginationRequest
	Status      vobj.GlobalStatus
	DashboardID uint32
}

// ListDashboardChartReply represents a reply to list dashboard charts
type ListDashboardChartReply struct {
	*PaginationReply
	Charts []*team.DashboardChart
}

// BatchUpdateDashboardStatusReq represents a request to batch update dashboard status
type BatchUpdateDashboardStatusReq struct {
	IDs    []uint32
	Status vobj.GlobalStatus
}

// BatchUpdateDashboardChartStatusReq represents a request to batch update dashboard chart status
type BatchUpdateDashboardChartStatusReq struct {
	IDs    []uint32
	Status vobj.GlobalStatus
}
