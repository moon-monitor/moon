package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameStrategy = "team_strategies"

type Strategy struct {
	do.TeamModel

	GroupID uint32              `gorm:"column:group_id;type:int unsigned;not null;comment:组ID" json:"group_id"`
	Name    string              `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark  string              `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status  vobj.StrategyStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`

	Group *StrategyGroup `gorm:"foreignKey:GroupID;references:ID" json:"group"`
}

func (s *Strategy) TableName() string {
	return tableNameStrategy
}
