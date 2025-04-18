package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToTeamAuditListRequest(req *palace.GetTeamAuditListRequest) *bo.TeamAuditListRequest {
	if validate.IsNil(req) {
		return nil
	}
	return &bo.TeamAuditListRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Status:            slices.Map(req.GetStatus(), func(status common.TeamAuditStatus) vobj.StatusAudit { return vobj.StatusAudit(status) }),
		Keyword:           req.GetKeyword(),
		Actions:           slices.Map(req.GetActions(), func(action common.TeamAuditAction) vobj.AuditAction { return vobj.AuditAction(action) }),
		UserID:            req.GetUserId(),
	}
}

func ToTeamAuditItem(audit do.TeamAudit) *common.TeamAuditItem {
	if validate.IsNil(audit) {
		return nil
	}
	return &common.TeamAuditItem{
		Id:        audit.GetID(),
		User:      UserToUserBaseItemProto(audit.GetCreator()),
		Status:    common.TeamAuditStatus(audit.GetStatus()),
		Reason:    audit.GetReason(),
		CreatedAt: audit.GetCreatedAt().Format(time.DateTime),
		Team:      ToTeamBaseItem(audit.GetTeam()),
		Action:    common.TeamAuditAction(audit.GetAction()),
	}
}

func ToTeamAuditItems(audits []do.TeamAudit) []*common.TeamAuditItem {
	return slices.Map(audits, ToTeamAuditItem)
}
