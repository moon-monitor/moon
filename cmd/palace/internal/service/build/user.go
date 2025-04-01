package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
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
