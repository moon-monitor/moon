package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameDashboard = "team_dashboards"

type Dashboard struct {
	do.TeamModel
	Title    string            `gorm:"column:title;type:varchar(255);not null;comment:标题" json:"title"`
	Remark   string            `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	Status   vobj.GlobalStatus `gorm:"column:status;type:tinyint;not null;default:0;comment:状态" json:"status"`
	ColorHex string            `gorm:"column:color_hex;type:varchar(20);not null;comment:颜色Hex" json:"colorHex"`
	Charts   []*DashboardChart `gorm:"foreignKey:DashboardID;references:ID" json:"charts"`
}

func (d *Dashboard) TableName() string {
	return tableNameDashboard
}
