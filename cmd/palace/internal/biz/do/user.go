package do

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

type User interface {
	Base
	GetUsername() string
	GetNickname() string
	GetEmail() crypto.String
	GetPhone() crypto.String
	GetPassword() string
	GetSalt() string
	GetGender() vobj.Gender
	GetAvatar() string
	GetStatus() vobj.UserStatus
	GetPosition() vobj.Role
	GetRoles() []Role
	GetTeams() []Team
	ValidatePassword(p string) bool
}
