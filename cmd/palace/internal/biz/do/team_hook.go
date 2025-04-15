package do

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/kv"
)

type NoticeHook interface {
	GetHookID() uint32
	GetTeamID() uint32
	GetName() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetURL() string
	GetMethod() vobj.HTTPMethod
	GetSecret() string
	GetHeaders() kv.StringMap
	GetApp() vobj.HookApp
	GetNoticeGroups() []NoticeGroup
}

type NoticeGroup interface {
	GetID() uint32
	GetTeamID() uint32
	GetName() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetHooks() []NoticeHook
	GetNoticeMembers() []NoticeMember
}

type NoticeMember interface {
	GetID() uint32
	GetTeamID() uint32
	GetNoticeGroupID() uint32
	GetMemberID() uint32
	GetNoticeType() vobj.NoticeType
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetNoticeGroup() NoticeGroup
	GetMember() TeamMember
}
