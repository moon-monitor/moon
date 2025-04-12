package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameOAuthUser = "sys_oauth_users"

type UserOAuth struct {
	do.BaseModel
	OAuthID uint32        `gorm:"column:oauth_id;type:int unsigned;index:uk__oauth_id__sys_user_id__app,unique" json:"oauth_id"`
	UserID  uint32        `gorm:"column:user_id;type:int unsigned;not null;comment:关联用户id;index:uk__oauth_id__sys_user_id__app,unique" json:"user_id"`
	Row     string        `gorm:"column:row;type:text;comment:用户信息" json:"row"`
	APP     vobj.OAuthAPP `gorm:"column:app;type:tinyint;not null;comment:oauth应用;index:uk__oauth_id__sys_user_id__app,unique" json:"app"`
	User    *User         `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (s *UserOAuth) TableName() string {
	return tableNameOAuthUser
}
