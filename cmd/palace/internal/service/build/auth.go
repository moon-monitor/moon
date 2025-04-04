package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

func LoginSignToUserBaseProto(b *bo.LoginSign) *common.UserBaseItem {
	if b == nil || b.Base == nil {
		return nil
	}
	return &common.UserBaseItem{
		Username: b.Base.Username,
		Nickname: b.Base.Nickname,
		Avatar:   b.Base.Avatar,
		Gender:   common.Gender(b.Base.Gender),
		UserId:   b.Base.UserID,
	}
}

func LoginReply(b *bo.LoginSign) *palace.LoginReply {
	return &palace.LoginReply{
		Token:          b.Token,
		ExpiredSeconds: int32(b.ExpiredSeconds),
		User:           LoginSignToUserBaseProto(b),
	}
}
