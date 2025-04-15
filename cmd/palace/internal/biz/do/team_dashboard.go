package do

import (
	"time"

	"gorm.io/plugin/soft_delete"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type Dashboard interface {
	GetDashboardID() uint32
	GetTeamID() uint32
	GetTitle() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetColorHex() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() soft_delete.DeletedAt
	GetCharts() []DashboardChart
}

type DashboardChart interface {
	GetDashboardChartID() uint32
	GetTeamID() uint32
	GetDashboardID() uint32
	GetTitle() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetConfig() string
	GetSort() int32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() soft_delete.DeletedAt
	GetDashboard() Dashboard
}
