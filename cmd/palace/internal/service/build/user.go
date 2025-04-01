package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

// UserToUserItem converts a system.User to a common.UserItem
func UserToUserItemProto(user *system.User) *common.UserItem {
	if user == nil {
		return nil
	}

	return &common.UserItem{
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Gender:   common.Gender(user.Gender),
	}
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
