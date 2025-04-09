package do

import (
	"github.com/moon-monitor/moon/pkg/plugin/email"
)

var _ email.Config = (*EmailConfig)(nil)

type EmailConfig struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Host   string `json:"host"`
	Port   uint32 `json:"port"`
	Enable bool   `json:"enable"`
	Name   string `json:"name"`
}

func (e *EmailConfig) GetName() string {
	return e.Name
}

func (e *EmailConfig) GetUser() string {
	return e.User
}

func (e *EmailConfig) GetPass() string {
	return e.Pass
}

func (e *EmailConfig) GetHost() string {
	return e.Host
}

func (e *EmailConfig) GetPort() uint32 {
	return e.Port
}

func (e *EmailConfig) GetEnable() bool {
	return e.Enable
}
