package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

// ToSaveTeamNoticeHookRequest 转换保存钩子请求
func ToSaveTeamNoticeHookRequest(req *palace.SaveTeamNoticeHookRequest) *bo.SaveTeamNoticeHookRequest {
	if req == nil {
		return nil
	}
	return &bo.SaveTeamNoticeHookRequest{
		HookID:  req.GetHookID(),
		Name:    req.GetName(),
		Remark:  req.GetRemark(),
		Status:  vobj.GlobalStatus(req.GetStatus()),
		URL:     req.GetUrl(),
		Method:  vobj.HTTPMethod(req.GetMethod()),
		Secret:  req.GetSecret(),
		Headers: req.GetHeaders(),
		APP:     vobj.HookApp(req.GetApp()),
	}
}

// ToListTeamNoticeHookRequest 转换列表请求
func ToListTeamNoticeHookRequest(req *palace.ListTeamNoticeHookRequest) *bo.ListTeamNoticeHookRequest {
	if req == nil {
		return nil
	}
	return &bo.ListTeamNoticeHookRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Status:            vobj.GlobalStatus(req.GetStatus()),
		Keyword:           req.GetKeyword(),
		Apps:              slices.Map(req.GetApps(), func(app common.HookAPP) vobj.HookApp { return vobj.HookApp(app) }),
	}
}

// ToNoticeHookItem 转换钩子信息
func ToNoticeHookItem(hook do.NoticeHook) *common.NoticeHookItem {
	if hook == nil {
		return nil
	}
	return &common.NoticeHookItem{
		NoticeHookID: hook.GetID(),
		Name:         hook.GetName(),
		Remark:       hook.GetRemark(),
		Status:       common.GlobalStatus(hook.GetStatus()),
		Url:          hook.GetURL(),
		Method:       common.HTTPMethod(hook.GetMethod()),
		Secret:       hook.GetSecret(),
		Headers:      hook.GetHeaders(),
		CreatedAt:    hook.GetCreatedAt().Format(time.DateTime),
		UpdatedAt:    hook.GetUpdatedAt().Format(time.DateTime),
	}
}

// ToNoticeHookItems 转换钩子信息列表
func ToNoticeHookItems(hooks []do.NoticeHook) []*common.NoticeHookItem {
	return slices.Map(hooks, ToNoticeHookItem)
}
