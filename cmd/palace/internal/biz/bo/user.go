package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type UserUpdateInfo struct {
	do.User
	UserID   uint32
	Nickname string
	Avatar   string
	Gender   vobj.Gender
}

func (u *UserUpdateInfo) WithUser(user do.User) *UserUpdateInfo {
	u.User = user
	return u
}

func (u *UserUpdateInfo) GetUserID() uint32 {
	if u == nil {
		return 0
	}
	if u.User == nil {
		return u.UserID
	}
	return u.User.GetID()
}

func (u *UserUpdateInfo) GetNickname() string {
	if u == nil {
		return ""
	}
	return u.Nickname
}

func (u *UserUpdateInfo) GetAvatar() string {
	if u == nil {
		return ""
	}
	return u.Avatar
}

func (u *UserUpdateInfo) GetGender() vobj.Gender {
	if u == nil {
		return vobj.GenderUnknown
	}
	return u.Gender
}

type PasswordUpdateInfo struct {
	OldPassword  string
	NewPassword  string
	SendEmailFun SendEmailFun
}

type UpdateUserPasswordInfo struct {
	UserID         uint32
	Password       string
	Salt           string
	OriginPassword string
	SendEmailFun   SendEmailFun
}

type UpdateUserStatusRequest struct {
	UserIds []uint32
	Status  vobj.UserStatus
}

type ResetUserPasswordRequest struct {
	UserId       uint32
	SendEmailFun SendEmailFun
}

type UpdateUserPositionRequest struct {
	UserId   uint32
	Position vobj.Role
}
