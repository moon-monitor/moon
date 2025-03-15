package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

const tableNameDict = "team_dict"

type Dict struct {
	do.TeamModel
	Key   string `gorm:"column:key;type:varchar(64);not null;comment:字典key" json:"key"`
	Value string `gorm:"column:value;type:varchar(255);not null;comment:字典value" json:"value"`
}

func (u *Dict) TableName() string {
	return tableNameDict
}
