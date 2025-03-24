package system

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/crypto"
	"github.com/moon-monitor/moon/pkg/util/password"
)

const tableNameUser = "sys_users"

type User struct {
	do.BaseModel
	Username string          `gorm:"column:username;type:varchar(64);not null;index:idx__sys_user__username,priority:1;comment:用户名" json:"username"`
	Nickname string          `gorm:"column:nickname;type:varchar(64);not null;comment:昵称" json:"nickname"`
	Password string          `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"-"`
	Email    crypto.String   `gorm:"column:email;type:varchar(255);not null;comment:邮箱;uniqueIndex:uk__sys_user__email,priority:1;" json:"email"`
	Phone    crypto.String   `gorm:"column:phone;type:varchar(255);not null;index:idx__sys_user__phone,priority:1;comment:手机号" json:"phone"`
	Remark   string          `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Avatar   string          `gorm:"column:avatar;type:varchar(255);not null;comment:头像" json:"avatar"`
	Salt     string          `gorm:"column:salt;type:varchar(16);not null;comment:盐" json:"-"`
	Gender   vobj.Gender     `gorm:"column:gender;type:tinyint(2);not null;comment:性别" json:"gender"`
	Position vobj.Role       `gorm:"column:position;type:tinyint(2);not null;comment:系统默认角色类型" json:"position"`
	Status   vobj.UserStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`

	Roles []*SysRole `gorm:"many2many:sys_user_roles;foreignKey:ID;joinForeignKey:UserID;references:ID;joinReferences:RoleID" json:"roles"`
}

func (u *User) TableName() string {
	return tableNameUser
}

// ValidatePassword validate password
func (u *User) ValidatePassword(p string) bool {
	validate := password.New(p, u.Salt)
	return validate.EQ(u.Password)
}
