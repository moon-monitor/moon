package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type Dashboard interface {
	GetDashboardID() uint32
	GetTitle() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetColorHex() string
}

// SaveDashboardReq represents a request to save a dashboard
type SaveDashboardReq struct {
	dashboard Dashboard
	ID        uint32
	Title     string
	Remark    string
	Status    vobj.GlobalStatus
	ColorHex  string
}

func (d *SaveDashboardReq) GetDashboardID() uint32 {
	if d == nil || d.dashboard == nil {
		return 0
	}
	return d.dashboard.GetDashboardID()
}

func (d *SaveDashboardReq) GetTitle() string {
	return d.Title
}

func (d *SaveDashboardReq) GetRemark() string {
	return d.Remark
}

func (d *SaveDashboardReq) GetStatus() vobj.GlobalStatus {
	return d.Status
}

func (d *SaveDashboardReq) GetColorHex() string {
	return d.ColorHex
}

func (d *SaveDashboardReq) WithDashboard(dashboard Dashboard) Dashboard {
	d.dashboard = dashboard
	return d
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
