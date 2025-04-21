package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToBaseModel(model do.Base) do.BaseModel {
	if validate.IsNil(model) {
		return do.BaseModel{}
	}
	return do.BaseModel{
		ID:        model.GetID(),
		CreatedAt: model.GetCreatedAt(),
		UpdatedAt: model.GetUpdatedAt(),
		DeletedAt: model.GetDeletedAt(),
	}
}

func ToCreatorModel(model do.Creator) do.CreatorModel {
	if validate.IsNil(model) {
		return do.CreatorModel{}
	}
	return do.CreatorModel{
		BaseModel: ToBaseModel(model),
		CreatorID: model.GetCreatorID(),
	}
}

func ToTeamModel(model do.TeamBase) do.TeamModel {
	if validate.IsNil(model) {
		return do.TeamModel{}
	}
	return do.TeamModel{
		CreatorModel: ToCreatorModel(model),
		TeamID:       model.GetTeamID(),
	}
}
