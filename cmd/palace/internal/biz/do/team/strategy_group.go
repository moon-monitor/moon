package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameStrategyGroup = "team_strategy_group"

type StrategyGroup struct {
	do.TeamModel

	Name   string                   `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark string                   `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status vobj.StrategyGroupStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
}

func (s *StrategyGroup) TableName() string {
	return tableNameStrategyGroup
}
