package bo

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
