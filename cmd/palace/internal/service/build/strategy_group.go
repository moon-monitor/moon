package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToTeamStrategyGroupItem(group do.StrategyGroup) *common.TeamStrategyGroupItem {
	return &common.TeamStrategyGroupItem{
		Name:                group.GetName(),
		Remark:              group.GetRemark(),
		GroupId:             group.GetID(),
		Status:              common.GlobalStatus(group.GetStatus().GetValue()),
		StrategyCount:       0,
		EnableStrategyCount: 0,
		CreatedAt:           group.GetCreatedAt().Format(time.DateTime),
		UpdatedAt:           group.GetUpdatedAt().Format(time.DateTime),
		Creator:             ToUserBaseItem(group.GetCreator()),
	}
}

func ToTeamStrategyGroupItems(groups []do.StrategyGroup) []*common.TeamStrategyGroupItem {
	return slices.Map(groups, ToTeamStrategyGroupItem)
}
