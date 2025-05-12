package bo

import (
	"github.com/moon-monitor/moon/pkg/api/common"
	apicommon "github.com/moon-monitor/moon/pkg/api/rabbit/common"
)

type NoticeGroup interface {
	GetName() string
	GetSms() []string
	GetEmails() []string
	GetHooks() []string
	GetTemplates() map[common.NoticeType]string
	GetTemplate(noticeType common.NoticeType) string
	GetSmsTemplate() string
	GetEmailTemplate() string
	GetHookTemplate(app apicommon.HookAPP) string
}

func NewNoticeGroup(opts ...NoticeGroupOption) NoticeGroup {
	noticeGroup := &noticeGroup{
		templates: make(map[common.NoticeType]string, 7),
	}
	for _, opt := range opts {
		opt(noticeGroup)
	}
	return noticeGroup
}

type NoticeGroupOption func(noticeGroup *noticeGroup)

type noticeGroup struct {
	name      string
	sms       []string
	emails    []string
	hooks     []string
	templates map[common.NoticeType]string
}

// GetEmails implements NoticeGroup.
func (n *noticeGroup) GetEmails() []string {
	return n.emails
}

// GetHooks implements NoticeGroup.
func (n *noticeGroup) GetHooks() []string {
	return n.hooks
}

// GetName implements NoticeGroup.
func (n *noticeGroup) GetName() string {
	return n.name
}

// GetSms implements NoticeGroup.
func (n *noticeGroup) GetSms() []string {
	return n.sms
}

// GetTemplates implements NoticeGroup.
func (n *noticeGroup) GetTemplates() map[common.NoticeType]string {
	return n.templates
}

func (n *noticeGroup) GetTemplate(noticeType common.NoticeType) string {
	return n.templates[noticeType]
}

func (n *noticeGroup) GetSmsTemplate() string {
	return n.templates[common.NoticeType_NOTICE_TYPE_SMS]
}

func (n *noticeGroup) GetEmailTemplate() string {
	return n.templates[common.NoticeType_NOTICE_TYPE_EMAIL]
}

func (n *noticeGroup) GetHookTemplate(app apicommon.HookAPP) string {
	var template string
	switch app {
	case apicommon.HookAPP_DINGTALK:
		template = n.templates[common.NoticeType_NOTICE_TYPE_HOOK_DINGTALK]
	case apicommon.HookAPP_WECHAT:
		template = n.templates[common.NoticeType_NOTICE_TYPE_HOOK_WECHAT]
	case apicommon.HookAPP_FEISHU:
		template = n.templates[common.NoticeType_NOTICE_TYPE_HOOK_FEISHU]
	case apicommon.HookAPP_OTHER:
		template = n.templates[common.NoticeType_NOTICE_TYPE_HOOK_WEBHOOK]
	}
	return template
}

func WithNoticeGroupOptionName(name string) NoticeGroupOption {
	return func(noticeGroup *noticeGroup) {
		noticeGroup.name = name
	}
}

func WithNoticeGroupOptionSms(sms []string) NoticeGroupOption {
	return func(noticeGroup *noticeGroup) {
		noticeGroup.sms = sms
	}
}

func WithNoticeGroupOptionEmails(emails []string) NoticeGroupOption {
	return func(noticeGroup *noticeGroup) {
		noticeGroup.emails = emails
	}
}

func WithNoticeGroupOptionHooks(hooks []string) NoticeGroupOption {
	return func(noticeGroup *noticeGroup) {
		noticeGroup.hooks = hooks
	}
}

type Template interface {
	GetType() common.NoticeType
	GetTemplate() string
}

func WithNoticeGroupOptionTemplates(templates []Template) NoticeGroupOption {
	return func(noticeGroup *noticeGroup) {
		for _, template := range templates {
			noticeGroup.templates[template.GetType()] = template.GetTemplate()
		}
	}
}

func WithNoticeGroupOptionTemplate(template Template) NoticeGroupOption {
	return func(noticeGroup *noticeGroup) {
		noticeGroup.templates[template.GetType()] = template.GetTemplate()
	}
}
