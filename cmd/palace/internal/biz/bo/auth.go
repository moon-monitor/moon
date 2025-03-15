package bo

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/middleware"
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
	Base           *middleware.JwtBaseInfo
	Token          string
	ExpiredSeconds int64
}
