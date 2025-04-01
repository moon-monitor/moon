package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

// TeamToTeamItemProto 将系统Team对象转换为TeamItem proto对象
func TeamToTeamItemProto(team *system.Team) *common.TeamItem {
	if team == nil {
		return nil
	}

	teamItem := &common.TeamItem{
		Id:        team.ID,
		Uuid:      team.UUID.String(),
		Name:      team.Name,
		Remark:    team.Remark,
		Logo:      team.Logo,
		Status:    common.TeamStatus(team.Status),
		CreatedAt: team.CreatedAt.Format(time.RFC3339),
		UpdatedAt: team.UpdatedAt.Format(time.RFC3339),
	}

	// 添加领导者信息
	if team.Leader != nil {
		teamItem.Leader = UserToUserItemProto(team.Leader)
	}

	return teamItem
}

// TeamsToTeamItemProtos 将系统Team对象切片转换为TeamItem proto对象切片
func TeamsToTeamItemProtos(teams []*system.Team) []*common.TeamItem {
	if teams == nil {
		return nil
	}

	items := make([]*common.TeamItem, 0, len(teams))
	for _, team := range teams {
		items = append(items, TeamToTeamItemProto(team))
	}

	return items
}
