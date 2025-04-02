package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const (
	tableNameDashboard      = "team_dashboard"
	tableNameDashboardChart = "team_dashboard_chart"
)

type Dashboard struct {
	do.TeamModel
	Title    string            `gorm:"column:title;type:varchar(255);not null;comment:标题" json:"title"`
	Remark   string            `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	Status   vobj.GlobalStatus `gorm:"column:status;type:tinyint;not null;default:0;comment:状态" json:"status"`
	ColorHex string            `gorm:"column:color_hex;type:varchar(20);not null;comment:颜色Hex" json:"color_hex"`

	Charts []*DashboardChart `gorm:"foreignKey:DashboardID;references:ID" json:"charts"`
}

func (d *Dashboard) TableName() string {
	return tableNameDashboard
}

type DashboardChart struct {
	do.TeamModel
	DashboardID uint32            `gorm:"column:dashboard_id;type:int;not null;comment:仪表盘ID" json:"dashboard_id"`
	Title       string            `gorm:"column:title;type:varchar(255);not null;comment:标题" json:"title"`
	Remark      string            `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	Status      vobj.GlobalStatus `gorm:"column:status;type:tinyint;not null;default:0;comment:状态" json:"status"`
	Url         string            `gorm:"column:url;type:varchar(255);not null;comment:URL" json:"url"`
	Width       string            `gorm:"column:width;type:varchar(255);not null;comment:宽度" json:"width"`
	Height      string            `gorm:"column:height;type:varchar(255);not null;comment:高度" json:"height"`

	Dashboard *Dashboard `gorm:"foreignKey:DashboardID;references:ID" json:"dashboard"`
}

func (c *DashboardChart) TableName() string {
	return tableNameDashboardChart
}
