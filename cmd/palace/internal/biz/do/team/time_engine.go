package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameTimeEngine = "team_time_engines"

type TimeEngine struct {
	do.TeamModel

	Name   string                `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark string                `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status vobj.TimeEngineStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Rules  []*TimeEngineRule     `gorm:"many2many:team_time_engine__time_rules" json:"rules"`
}

func (t *TimeEngine) TableName() string {
	return tableNameTimeEngine
}
