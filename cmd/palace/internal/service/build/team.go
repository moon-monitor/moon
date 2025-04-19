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

// TeamToTeamItemProto 将系统Team对象转换为TeamItem proto对象
func TeamToTeamItemProto(team do.Team) *common.TeamItem {
	if validate.IsNil(team) {
		return nil
	}

	return &common.TeamItem{
		Id:              team.GetID(),
		Uuid:            team.GetUUID().String(),
		Name:            team.GetName(),
		Remark:          team.GetRemark(),
		Logo:            team.GetLogo(),
		Status:          common.TeamStatus(team.GetStatus()),
		Creator:         UserToUserBaseItemProto(team.GetCreator()),
		Leader:          UserToUserBaseItemProto(team.GetLeader()),
		Admins:          UsersToUserBaseItemsProto(team.GetAdmins()),
		CreatedAt:       team.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:       team.GetUpdatedAt().Format(time.RFC3339),
		MemberCount:     0,
		StrategyCount:   0,
		DatasourceCount: 0,
	}
}

// TeamsToTeamItemsProto 将系统Team对象切片转换为TeamItem proto对象切片
func TeamsToTeamItemsProto(teams []do.Team) []*common.TeamItem {
	if len(teams) == 0 {
		return nil
	}

	items := make([]*common.TeamItem, 0, len(teams))
	for _, team := range teams {
		items = append(items, TeamToTeamItemProto(team))
	}

	return items
}

func ToTeamListRequest(req *palace.GetTeamListRequest) *bo.TeamListRequest {
	return &bo.TeamListRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Keyword:           req.GetKeyword(),
		Status:            slices.Map(req.GetStatus(), func(status common.TeamStatus) vobj.TeamStatus { return vobj.TeamStatus(status) }),
		UserIds:           nil,
		LeaderId:          req.GetLeaderId(),
		CreatorId:         req.GetCreatorId(),
	}
}
