package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToUser(userDo do.User) *system.User {
	if validate.IsNil(userDo) {
		return nil
	}
	user, ok := userDo.(*system.User)
	if ok {
		return user
	}
	return &system.User{
		BaseModel: ToBaseModel(userDo),
		Username:  userDo.GetUsername(),
		Nickname:  userDo.GetNickname(),
		Email:     userDo.GetEmail(),
		Phone:     userDo.GetPhone(),
		Remark:    userDo.GetRemark(),
		Avatar:    userDo.GetAvatar(),
		Gender:    userDo.GetGender(),
		Position:  userDo.GetPosition(),
		Status:    userDo.GetStatus(),
		Roles:     ToRoles(userDo.GetRoles()),
		Teams:     ToTeams(userDo.GetTeams()),
	}
}

func ToUsers(userDos []do.User) []*system.User {
	return slices.Map(userDos, ToUser)
}
