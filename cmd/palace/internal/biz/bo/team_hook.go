package bo

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/kv"
)

// SaveTeamNoticeHookRequest 保存团队通知钩子请求
type SaveTeamNoticeHookRequest struct {
	hookDo  do.NoticeHook
	HookID  uint32            `json:"hookId"`
	Name    string            `json:"name"`
	Remark  string            `json:"remark"`
	Status  vobj.GlobalStatus `json:"status"`
	URL     string            `json:"url"`
	Method  vobj.HTTPMethod   `json:"method"`
	Secret  string            `json:"secret"`
	Headers kv.StringMap      `json:"headers"`
	APP     vobj.HookApp      `json:"app"`
}

func (r *SaveTeamNoticeHookRequest) GetHookID() uint32 {
	if r == nil || r.hookDo == nil {
		return 0
	}
	return r.hookDo.GetHookID()
}

func (r *SaveTeamNoticeHookRequest) GetTeamID() uint32 {
	if r == nil || r.hookDo == nil {
		return 0
	}
	return r.hookDo.GetTeamID()
}

func (r *SaveTeamNoticeHookRequest) GetCreatedAt() time.Time {
	if r == nil || r.hookDo == nil {
		return time.Now()
	}
	return r.hookDo.GetCreatedAt()
}

func (r *SaveTeamNoticeHookRequest) GetUpdatedAt() time.Time {
	if r == nil || r.hookDo == nil {
		return time.Now()
	}
	return r.hookDo.GetUpdatedAt()
}

func (r *SaveTeamNoticeHookRequest) GetNoticeGroups() []do.NoticeGroup {
	if r == nil || r.hookDo == nil {
		return nil
	}
	return r.hookDo.GetNoticeGroups()
}

func (r *SaveTeamNoticeHookRequest) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}

func (r *SaveTeamNoticeHookRequest) GetRemark() string {
	if r == nil {
		return ""
	}
	return r.Remark
}

func (r *SaveTeamNoticeHookRequest) GetStatus() vobj.GlobalStatus {
	if r == nil {
		return vobj.GlobalStatusUnknown
	}
	return r.Status
}

func (r *SaveTeamNoticeHookRequest) GetURL() string {
	if r == nil {
		return ""
	}
	return r.URL
}

func (r *SaveTeamNoticeHookRequest) GetMethod() vobj.HTTPMethod {
	if r == nil {
		return vobj.HTTPMethodPost
	}
	return r.Method
}

func (r *SaveTeamNoticeHookRequest) GetSecret() string {
	if r == nil {
		return ""
	}
	return r.Secret
}

func (r *SaveTeamNoticeHookRequest) GetHeaders() kv.StringMap {
	if r == nil {
		return nil
	}
	return r.Headers
}

func (r *SaveTeamNoticeHookRequest) GetApp() vobj.HookApp {
	if r == nil {
		return vobj.HookAppOther
	}
	return r.APP
}

func (r *SaveTeamNoticeHookRequest) WithUpdateHookRequest(hook do.NoticeHook) do.NoticeHook {
	r.hookDo = hook
	return r
}

// ListTeamNoticeHookRequest 列表请求
type ListTeamNoticeHookRequest struct {
	*PaginationRequest
	Status  vobj.GlobalStatus `json:"status"`
	Keyword string            `json:"keyword"`
	Apps    []vobj.HookApp    `json:"apps"`
}

func (r *ListTeamNoticeHookRequest) ToListTeamNoticeHookReply(hooks []do.NoticeHook) *ListTeamNoticeHookReply {
	return &ListTeamNoticeHookReply{
		PaginationReply: r.ToReply(),
		Hooks:           hooks,
	}
}

// ListTeamNoticeHookReply 列表响应
type ListTeamNoticeHookReply struct {
	*PaginationReply
	Hooks []do.NoticeHook `json:"hooks"`
}

// UpdateTeamNoticeHookStatusRequest 更新状态请求
type UpdateTeamNoticeHookStatusRequest struct {
	HookID uint32            `json:"hookId"`
	Status vobj.GlobalStatus `json:"status"`
}
