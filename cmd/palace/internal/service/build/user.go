package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

// ToUserItem converts a system.User to a common.UserItem
func ToUserItem(user do.User) *common.UserItem {
	if validate.IsNil(user) {
		return nil
	}

	return &common.UserItem{
		Username:  user.GetUsername(),
		Nickname:  user.GetNickname(),
		Avatar:    user.GetAvatar(),
		Gender:    common.Gender(user.GetGender()),
		Email:     string(user.GetEmail()),
		Phone:     string(user.GetPhone()),
		Remark:    user.GetRemark(),
		Position:  common.UserPosition(user.GetPosition()),
		Status:    common.UserStatus(user.GetStatus()),
		CreatedAt: user.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: user.GetUpdatedAt().Format(time.DateTime),
		UserID:    user.GetID(),
	}
}

func ToUserBaseItem(user do.User) *common.UserBaseItem {
	if validate.IsNil(user) {
		return nil
	}

	return &common.UserBaseItem{
		Username: user.GetUsername(),
		Nickname: user.GetNickname(),
		Avatar:   user.GetAvatar(),
		Gender:   common.Gender(user.GetGender()),
		UserID:   user.GetID(),
	}
}

// ToUserItems converts a slice of system.User to a slice of common.UserItem
func ToUserItems(users []do.User) []*common.UserItem {
	return slices.Map(users, ToUserItem)
}

// ToUserBaseItems converts a slice of system.User to a slice of common.UserBaseItem
func ToUserBaseItems(users []do.User) []*common.UserBaseItem {
	return slices.Map(users, ToUserBaseItem)
}

func ToSelfUpdateInfo(req *palacev1.UpdateSelfInfoRequest) *bo.UserUpdateInfo {
	if req == nil {
		return nil
	}

	return &bo.UserUpdateInfo{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Gender:   vobj.Gender(req.Gender),
	}
}

func ToUserUpdateInfo(req *palacev1.UpdateUserRequest) *bo.UserUpdateInfo {
	if req == nil {
		return nil
	}

	return &bo.UserUpdateInfo{
		UserID:   req.GetUserId(),
		Nickname: req.GetNickname(),
		Avatar:   req.GetAvatar(),
		Gender:   vobj.Gender(req.GetGender()),
	}
}

// ToPasswordUpdateInfo converts an API password update request to a business object
func ToPasswordUpdateInfo(req *palacev1.UpdateSelfPasswordRequest, sendEmailFun bo.SendEmailFun) *bo.PasswordUpdateInfo {
	if req == nil {
		return nil
	}

	return &bo.PasswordUpdateInfo{
		OldPassword:  req.OldPassword,
		NewPassword:  req.NewPassword,
		SendEmailFun: sendEmailFun,
	}
}

func ToUserListRequest(req *palacev1.GetUserListRequest) *bo.UserListRequest {
	return &bo.UserListRequest{
		PaginationRequest: ToPaginationRequest(req.GetPagination()),
		Status:            slices.Map(req.GetStatus(), func(status common.UserStatus) vobj.UserStatus { return vobj.UserStatus(status) }),
		Position:          slices.Map(req.GetPosition(), func(position common.UserPosition) vobj.Role { return vobj.Role(position) }),
		Keyword:           req.GetKeyword(),
	}
}
