package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func LoginSignToUserBaseItem(b *bo.LoginSign) *common.UserBaseItem {
	if b == nil || b.Base == nil {
		return nil
	}
	return &common.UserBaseItem{
		Username: b.Base.Username,
		Nickname: b.Base.Nickname,
		Avatar:   b.Base.Avatar,
		Gender:   common.Gender(b.Base.Gender),
		UserID:   b.Base.UserID,
	}
}

func LoginReply(b *bo.LoginSign) *palace.LoginReply {
	return &palace.LoginReply{
		Token:          b.Token,
		ExpiredSeconds: int32(b.ExpiredSeconds),
		User:           LoginSignToUserBaseItem(b),
	}
}

func ToTLS(tls *common.TLS) *do.TLS {
	if validate.IsNil(tls) {
		return nil
	}
	return &do.TLS{
		ServerName: tls.GetServerName(),
		ClientCert: tls.GetClientCert(),
		ClientKey:  tls.GetClientKey(),
	}
}

func ToProtoTLS(tls *do.TLS) *common.TLS {
	if validate.IsNil(tls) {
		return nil
	}
	return &common.TLS{
		ServerName: tls.ServerName,
		ClientCert: tls.ClientCert,
		ClientKey:  tls.ClientKey,
	}
}

func ToBasicAuth(basicAuth *common.BasicAuth) *do.BasicAuth {
	if validate.IsNil(basicAuth) {
		return nil
	}
	return &do.BasicAuth{
		Username: basicAuth.GetUsername(),
		Password: basicAuth.GetPassword(),
	}
}

func ToProtoBasicAuth(basicAuth *do.BasicAuth) *common.BasicAuth {
	if validate.IsNil(basicAuth) {
		return nil
	}
	return &common.BasicAuth{
		Username: basicAuth.Username,
		Password: basicAuth.Password,
	}
}
