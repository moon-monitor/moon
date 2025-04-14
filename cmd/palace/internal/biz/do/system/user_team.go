package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

const tableNameUserTeam = "user_teams"

type UserTeam struct {
	do.BaseModel
	UserID uint32 `gorm:"column:user_id;not null;index:idx_user_id;int(10) unsigned;comment:用户ID" json:"userID"`
	TeamID uint32 `gorm:"column:team_id;not null;index:idx_team_id;int(10) unsigned;comment:团队ID" json:"teamID"`
	User   *User  `gorm:"foreignKey:UserID" json:"user"`
	Team   *Team  `gorm:"foreignKey:TeamID" json:"team"`
}

func (u *UserTeam) TableName() string {
	return tableNameUserTeam
}
