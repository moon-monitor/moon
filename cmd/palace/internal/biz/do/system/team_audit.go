package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameTeamAudit = "sys_team_audits"

type TeamAudit struct {
	do.CreatorModel

	TeamID uint32           `gorm:"column:team_id;type:int unsigned;not null;comment:团队ID" json:"team_id,omitempty"`
	Status vobj.StatusAudit `gorm:"column:status;type:tinyint(2);not null;comment:审批状态" json:"status"`
	Action vobj.AuditAction `gorm:"column:action;type:tinyint(2);not null;comment:操作" json:"action"`
	Reason string           `gorm:"column:reason;type:varchar(255);not null;comment:原因" json:"reason"`

	User *User `gorm:"foreignKey:CreatorID;references:ID" json:"user"`
	Team *Team `gorm:"foreignKey:TeamID;references:ID" json:"team"`
}

func (u *TeamAudit) TableName() string {
	return tableNameTeamAudit
}
