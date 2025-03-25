package bo

import (
	"encoding/json"
	"strings"

	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/plugin/email"
	"github.com/moon-monitor/moon/pkg/util/template"
)

type SendEmailConfig struct {
	Email        string
	TemplateName string
	Parameters   string
	Subject      string
	ContentType  string
	Attachment   string
	Cc           string

	config *config.EmailConfig
	global map[string]*config.EmailConfig
}

func (s *SendEmailConfig) WithGlobalConfig(config map[string]*config.EmailConfig) *SendEmailConfig {
	s.global = config
	return s
}

func (s *SendEmailConfig) Send(msg []byte) error {
	body := make(map[string]any)
	_ = json.Unmarshal([]byte(s.Parameters), &body)
	email := email.New(s.config)
	email.SetTo(s.Email)
	email.SetSubject(template.HtmlFormatterX(s.Subject, body))
	email.SetBody(template.HtmlFormatterX(s.Parameters, body), s.ContentType)
	if s.Attachment != "" {
		email.SetAttach(template.HtmlFormatterX(s.Attachment, body))
	}
	if s.Cc != "" {
		cc := strings.Split(template.HtmlFormatterX(s.Cc, body), ",")
		email.SetCc(cc...)
	}
	return email.Send()
}
