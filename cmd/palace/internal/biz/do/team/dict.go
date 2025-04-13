package team

import (
	"time"

	"gorm.io/plugin/soft_delete"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameDict = "team_dictionaries"

type Dict struct {
	do.TeamModel
	Key      string            `gorm:"column:key;type:varchar(64);not null;comment:字典key" json:"key"`
	Value    string            `gorm:"column:value;type:varchar(255);not null;comment:字典value" json:"value"`
	Lang     string            `gorm:"column:lang;type:varchar(16);not null;comment:语言" json:"lang"`
	Color    string            `gorm:"column:color;type:varchar(16);not null;comment:颜色Hex" json:"color"`
	DictType vobj.DictType     `gorm:"column:type;type:tinyint(2);not null;comment:字典类型" json:"type"`
	Status   vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
}

func (u *Dict) GetTeamID() uint32 {
	if u == nil {
		return 0
	}
	return u.TeamID
}

func (u *Dict) GetID() uint32 {
	if u == nil {
		return 0
	}
	return u.ID
}

func (u *Dict) GetKey() string {
	if u == nil {
		return ""
	}
	return u.Key
}

func (u *Dict) GetValue() string {
	if u == nil {
		return ""
	}
	return u.Value
}

func (u *Dict) GetStatus() vobj.GlobalStatus {
	if u == nil {
		return vobj.GlobalStatusUnknown
	}
	return u.Status
}

func (u *Dict) GetType() vobj.DictType {
	if u == nil {
		return vobj.DictTypeUnknown
	}
	return u.DictType
}

func (u *Dict) GetColor() string {
	if u == nil {
		return ""
	}
	return u.Color
}

func (u *Dict) GetLang() string {
	if u == nil {
		return ""
	}
	return u.Lang
}

func (u *Dict) GetCreatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.CreatedAt
}

func (u *Dict) GetUpdatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.UpdatedAt
}

func (u *Dict) GetDeletedAt() soft_delete.DeletedAt {
	if u == nil {
		return 0
	}
	return u.DeletedAt
}

func (u *Dict) TableName() string {
	return tableNameDict
}
