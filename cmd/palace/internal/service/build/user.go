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

func ToUserUpdateInfo(userUpdateInfo *palacev1.UpdateSelfInfoRequest) *bo.UserUpdateInfo {
	if userUpdateInfo == nil {
		return nil
	}

	return &bo.UserUpdateInfo{
		Nickname: userUpdateInfo.Nickname,
		Avatar:   userUpdateInfo.Avatar,
		Gender:   vobj.Gender(userUpdateInfo.Gender),
	}
}
