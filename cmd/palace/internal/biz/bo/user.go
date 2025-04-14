package bo

import "github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"

type User interface {
}

type UserUpdateInfo struct {
	Nickname string
	Avatar   string
	Gender   vobj.Gender
}

type PasswordUpdateInfo struct {
	OldPassword string
	NewPassword string
}

type UpdateUserPasswordInfo struct {
	UserID   uint32
	Password string
	Salt     string
}
