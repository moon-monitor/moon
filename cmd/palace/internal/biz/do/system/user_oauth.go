package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameOAuthUser = "sys_oauth_user"

type OAuthUser struct {
	do.BaseModel
	OAuthID   uint32        `gorm:"column:oauth_id;type:int unsigned;index:uk__oauth_id__sys_user_id__app,unique" json:"oauth_id"`
	SysUserID uint32        `gorm:"column:sys_user_id;type:int unsigned;not null;comment:关联用户id;index:uk__oauth_id__sys_user_id__app,unique" json:"sys_user_id"`
	Row       string        `gorm:"column:row;type:text;comment:用户信息" json:"row"`
	APP       vobj.OAuthAPP `gorm:"column:app;type:tinyint;not null;comment:oauth应用;index:uk__oauth_id__sys_user_id__app,unique" json:"app"`

	User *User `gorm:"foreignKey:SysUserID;references:ID" json:"user"`
}

func (s *OAuthUser) TableName() string {
	return tableNameOAuthUser
}
