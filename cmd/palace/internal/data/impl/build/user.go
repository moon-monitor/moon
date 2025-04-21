package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToUser(userDo do.User) *system.User {
	if userDo == nil {
		return nil
	}
	user, ok := userDo.(*system.User)
	if ok {
		return user
	}
	return &system.User{
		BaseModel: do.BaseModel{
			ID:        userDo.GetID(),
			CreatedAt: userDo.GetCreatedAt(),
			UpdatedAt: userDo.GetUpdatedAt(),
		},
		Username: userDo.GetUsername(),
		Nickname: userDo.GetNickname(),
		Email:    userDo.GetEmail(),
		Phone:    userDo.GetPhone(),
		Remark:   userDo.GetRemark(),
		Avatar:   userDo.GetAvatar(),
		Gender:   userDo.GetGender(),
		Position: userDo.GetPosition(),
		Status:   userDo.GetStatus(),
		Roles:    ToRoles(userDo.GetRoles()),
		Teams:    ToTeams(userDo.GetTeams()),
	}
}

func ToUsers(userDos []do.User) []*system.User {
	return slices.Map(userDos, ToUser)
}
