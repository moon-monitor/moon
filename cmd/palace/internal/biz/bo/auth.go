package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
	"github.com/moon-monitor/moon/pkg/api/palace"
)

type Captcha struct {
	Id             string `json:"id"`
	B64s           string `json:"b64s"`
	Answer         string `json:"answer"`
	ExpiredSeconds int64  `json:"expired_seconds"`
}

type CaptchaVerify struct {
	Id     string `json:"id"`
	Answer string `json:"answer"`
	Clear  bool   `json:"clear"`
}

type LoginByPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSign struct {
	Base           *middleware.JwtBaseInfo `json:"base"`
	Token          string                  `json:"token"`
	ExpiredSeconds int64                   `json:"expired_seconds"`
}

// BaseToProto Convert BaseInfo to proto's BaseInfo.
func (b *LoginSign) BaseToProto() *palace.UserBaseItem {
	if b == nil || b.Base == nil {
		return nil
	}
	return &palace.UserBaseItem{
		Username: b.Base.Username,
		Nickname: b.Base.Nickname,
		Avatar:   b.Base.Avatar,
		Gender:   palace.Gender(b.Base.Gender),
		UserId:   b.Base.UserID,
	}
}

// LoginReply Reply of login.
func (b *LoginSign) LoginReply() *palace.LoginReply {
	return &palace.LoginReply{
		Token:          b.Token,
		ExpiredSeconds: b.ExpiredSeconds,
		User:           b.BaseToProto(),
	}
}
