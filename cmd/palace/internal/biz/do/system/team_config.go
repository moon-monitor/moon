package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

const tableNameTeamConfig = "sys_team_configs"

type TeamConfig struct {
	do.TeamModel

	Email crypto.Object[do.EmailConfig] `gorm:"column:email;type:text;not null;comment:邮件配置" json:"email"`
}

func (u *TeamConfig) TableName() string {
	return tableNameTeamConfig
}
