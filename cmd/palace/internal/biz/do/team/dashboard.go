package team

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"gorm.io/plugin/soft_delete"
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

func (d *Dashboard) GetDashboardID() uint32 {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetTeamID() uint32 {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetTitle() string {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetRemark() string {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetStatus() vobj.GlobalStatus {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetColorHex() string {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetCreatedAt() time.Time {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetUpdatedAt() time.Time {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetDeletedAt() soft_delete.DeletedAt {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) GetCharts() []do.DashboardChart {
	//TODO implement me
	panic("implement me")
}

func (d *Dashboard) TableName() string {
	return tableNameDashboard
}
