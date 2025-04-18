package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToDictProto(dictItem do.TeamDict) *common.TeamDictItem {
	return &common.TeamDictItem{
		TeamID:    dictItem.GetTeamID(),
		DictID:    dictItem.GetID(),
		CreatedAt: dictItem.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: dictItem.GetUpdatedAt().Format(time.DateTime),
		Key:       dictItem.GetKey(),
		Value:     dictItem.GetValue(),
		Lang:      dictItem.GetLang(),
		Color:     dictItem.GetColor(),
		DictType:  common.DictType(dictItem.GetType()),
		Status:    common.GlobalStatus(dictItem.GetStatus()),
	}
}

func ToDictProtos(dictItems []do.TeamDict) []*common.TeamDictItem {
	return slices.Map(dictItems, ToDictProto)
}
