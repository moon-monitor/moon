package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/plugin/email"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

var _ do.EmailConfig = (*EmailConfig)(nil)

const tableNameConfigEmail = "team_config_emails"

// EmailConfig represents email configuration for a team
type EmailConfig struct {
	do.TeamModel
	Name   string                    `gorm:"column:name;type:varchar(20);not null;comment:配置名称" json:"name"`
	Remark string                    `gorm:"column:remark;type:varchar(200);not null;comment:配置备注" json:"remark"`
	Status vobj.GlobalStatus         `gorm:"column:status;type:tinyint(2);not null;default:0;comment:状态" json:"status"`
	Email  *crypto.Object[*do.Email] `gorm:"column:email;type:text;not null;comment:邮件配置" json:"email"`
}

func (c *EmailConfig) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *EmailConfig) GetRemark() string {
	if c == nil {
		return ""
	}
	return c.Remark
}

func (c *EmailConfig) GetStatus() vobj.GlobalStatus {
	if c == nil {
		return vobj.GlobalStatusUnknown
	}
	return c.Status
}

func (c *EmailConfig) GetEmailConfig() email.Config {
	if c == nil || c.Email == nil {
		return nil
	}
	return c.Email.Get()
}

func (c *EmailConfig) TableName() string {
	return tableNameConfigEmail
}
