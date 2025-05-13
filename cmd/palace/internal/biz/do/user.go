package do

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/util/crypto"
)

type User interface {
	Base
	cache.Object
	GetUsername() string
	GetNickname() string
	GetEmail() crypto.String
	GetPhone() crypto.String
	GetRemark() string
	GetPassword() string
	GetSalt() string
	GetGender() vobj.Gender
	GetAvatar() string
	GetStatus() vobj.UserStatus
	GetPosition() vobj.Role
	GetRoles() []Role
	GetTeams() []Team
	ValidatePassword(p string) bool
	SetEmail(email crypto.String)
}

type UserOAuth interface {
	Base
	GetOpenID() string
	GetAPP() vobj.OAuthAPP
	GetUserID() uint32
	GetRow() string
	GetUser() User
	SetUser(user User)
}
