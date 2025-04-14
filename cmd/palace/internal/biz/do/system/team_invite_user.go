package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameTeamInviteUser = "team_invite_users"

type TeamInviteUser struct {
	do.CreatorModel
	TeamID       uint32    `gorm:"index:idx_team_invite_user_team_id;column:team_id;not null;type:int(10) unsigned;comment:团队ID" json:"teamID"`
	InviteUserID uint32    `gorm:"index:idx_team_invite_user_invite_user_id;column:invite_user_id;not null;type:int(10) unsigned;comment:被邀请用户ID" json:"inviteUserID"`
	InviteUser   *User     `gorm:"foreignKey:InviteUserID;references:ID" json:"inviteUser"`
	Position     vobj.Role `gorm:"column:position;type:tinyint(2);not null;comment:职位" json:"position"`
}

func (t *TeamInviteUser) TableName() string {
	return tableNameTeamInviteUser
}
