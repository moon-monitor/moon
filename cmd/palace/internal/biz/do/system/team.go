package system

import (
	"time"

	"github.com/google/uuid"
	
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameTeam = "sys_teams"

type Team struct {
	do.CreatorModel
	Name      string            `gorm:"column:name;type:varchar(64);not null;uniqueIndex:idx__team__name,priority:1;comment:团队空间名" json:"name"`
	Status    vobj.TeamStatus   `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Remark    string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Logo      string            `gorm:"column:logo;type:varchar(255);not null;comment:团队logo" json:"logo"`
	LeaderID  uint32            `gorm:"column:leader_id;type:int unsigned;not null;index:sys_teams__sys_users,priority:1;comment:负责人" json:"leader_id"`
	UUID      uuid.UUID         `gorm:"column:uuid;type:BINARY(16);not null" json:"uuid"`
	Capacity  vobj.TeamCapacity `gorm:"column:capacity;type:tinyint(2);not null;comment:团队容量(套餐)" json:"capacity"`
	Leader    *User             `gorm:"foreignKey:LeaderID;references:ID" json:"leader"`
	Admins    []*User           `gorm:"many2many:sys_team_admins" json:"admins"`
	Resources []*Resource       `gorm:"many2many:sys_team_resources" json:"resources"`
	DBName    string            `gorm:"column:db_name;type:varchar(64);not null;comment:数据库名" json:"db_name"`
}

func (u *Team) GetCreatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.CreatedAt
}

func (u *Team) GetUpdatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.UpdatedAt
}

func (u *Team) GetTeamID() uint32 {
	if u == nil {
		return 0
	}
	return u.ID
}

func (u *Team) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *Team) GetRemark() string {
	if u == nil {
		return ""
	}
	return u.Remark
}

func (u *Team) GetLogo() string {
	if u == nil {
		return ""
	}
	return u.Logo
}

func (u *Team) GetStatus() vobj.TeamStatus {
	if u == nil {
		return vobj.TeamStatusUnknown
	}
	return u.Status
}

func (u *Team) GetLeaderID() uint32 {
	if u == nil {
		return 0
	}
	return u.LeaderID
}

func (u *Team) GetUUID() uuid.UUID {
	if u == nil {
		return uuid.UUID{}
	}
	return u.UUID
}

func (u *Team) GetCapacity() vobj.TeamCapacity {
	if u == nil {
		return vobj.TeamCapacityDefault
	}
	return u.Capacity
}

func (u *Team) GetDBName() string {
	if u == nil {
		return ""
	}
	return u.DBName
}

func (u *Team) TableName() string {
	return tableNameTeam
}
