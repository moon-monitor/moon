package do

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/plugin/email"
)

var _ email.Config = (*Email)(nil)

type Email struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Host   string `json:"host"`
	Port   uint32 `json:"port"`
	Enable bool   `json:"enable"`
	Name   string `json:"name"`
}

func (e *Email) GetName() string {
	return e.Name
}

func (e *Email) GetUser() string {
	return e.User
}

func (e *Email) GetPass() string {
	return e.Pass
}

func (e *Email) GetHost() string {
	return e.Host
}

func (e *Email) GetPort() uint32 {
	return e.Port
}

func (e *Email) GetEnable() bool {
	return e.Enable
}

type EmailConfig interface {
	GetEmailConfigID() uint32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetName() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetEmailConfig() email.Config
}
