package do

import (
	"encoding/json"

	"github.com/moon-monitor/moon/pkg/api/common"
	apicommon "github.com/moon-monitor/moon/pkg/api/rabbit/common"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
)

var _ cache.Object = (*NoticeGroupConfig)(nil)

type NoticeGroupConfig struct {
	Name      string                       `json:"name"`
	Sms       []string                     `json:"sms"`
	Emails    []string                     `json:"emails"`
	Hooks     []string                     `json:"hooks"`
	Templates map[common.NoticeType]string `json:"templates"`
}

// GetEmailTemplate implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetEmailTemplate() string {
	if n == nil {
		return ""
	}
	return n.Templates[common.NoticeType_NOTICE_TYPE_EMAIL]
}

// GetHookTemplate implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetHookTemplate(app apicommon.HookAPP) string {
	if n == nil {
		return ""
	}
	return n.Templates[common.NoticeType_NOTICE_TYPE_HOOK_DINGTALK]
}

// GetSmsTemplate implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetSmsTemplate() string {
	if n == nil {
		return ""
	}
	return n.Templates[common.NoticeType_NOTICE_TYPE_SMS]
}

// GetTemplate implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetTemplate(noticeType common.NoticeType) string {
	if n == nil {
		return ""
	}
	return n.Templates[noticeType]
}

// GetEmails implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetEmails() []string {
	if n == nil {
		return nil
	}
	return n.Emails
}

// GetHooks implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetHooks() []string {
	if n == nil {
		return nil
	}
	return n.Hooks
}

// GetName implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetName() string {
	if n == nil {
		return ""
	}
	return n.Name
}

// GetSms implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetSms() []string {
	if n == nil {
		return nil
	}
	return n.Sms
}

// GetTemplates implements bo.NoticeGroup.
func (n *NoticeGroupConfig) GetTemplates() map[common.NoticeType]string {
	if n == nil {
		return nil
	}
	return n.Templates
}

// MarshalBinary implements cache.Object.
func (n *NoticeGroupConfig) MarshalBinary() (data []byte, err error) {
	return json.Marshal(n)
}

// UnmarshalBinary implements cache.Object.
func (n *NoticeGroupConfig) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *NoticeGroupConfig) UniqueKey() string {
	return n.Name
}
