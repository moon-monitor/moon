package do

import (
	"encoding/json"

	"github.com/moon-monitor/moon/pkg/api/common"
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
