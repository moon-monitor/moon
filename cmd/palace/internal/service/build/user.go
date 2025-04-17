package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

// UserToUserItemProto converts a system.User to a common.UserItem
func UserToUserItemProto(user do.User) *common.UserItem {
	if validate.IsNil(user) {
		return nil
	}

	return &common.UserItem{
		Username: user.GetUsername(),
		Nickname: user.GetNickname(),
		Avatar:   user.GetAvatar(),
		Gender:   common.Gender(user.GetGender()),
	}
}

// UsersToUserItemsProto converts a slice of system.User to a slice of common.UserItem
func UsersToUserItemsProto(users []do.User) []*common.UserItem {
	return slices.Map(users, UserToUserItemProto)
}

func ToUserUpdateInfo(req *palacev1.UpdateSelfInfoRequest) *bo.UserUpdateInfo {
	if req == nil {
		return nil
	}

	return &bo.UserUpdateInfo{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   vobj.Gender(req.Gender),
	}
}

// ToPasswordUpdateInfo converts an API password update request to a business object
func ToPasswordUpdateInfo(req *palacev1.UpdateSelfPasswordRequest) *bo.PasswordUpdateInfo {
	if req == nil {
		return nil
	}

	return &bo.PasswordUpdateInfo{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
}
