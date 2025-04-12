package do

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"

	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
)

type ORMModel interface {
	sql.Scanner
	driver.Valuer
}

// BaseModel gorm base model
type BaseModel struct {
	ctx context.Context `gorm:"-"`

	ID        uint32                `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time             `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time             `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at,omitempty"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint;not null;default:0;" json:"deleted_at,omitempty"`
}

// WithContext set context
func (u *BaseModel) WithContext(ctx context.Context) *BaseModel {
	u.ctx = ctx
	return u
}

// GetContext get context
func (u *BaseModel) GetContext() context.Context {
	return u.ctx
}

type CreatorModel struct {
	BaseModel
	CreatorID uint32 `gorm:"column:creator;type:int unsigned;not null;comment:创建者" json:"creator_id,omitempty"`
}

func (u *CreatorModel) BeforeCreate(tx *gorm.DB) (err error) {
	var exist bool
	u.CreatorID, exist = permission.GetUserIDByContext(u.GetContext())
	if !exist || u.CreatorID == 0 {
		return merr.ErrorInternalServerError("user id not found")
	}
	tx.WithContext(u.GetContext())
	return
}

type TeamModel struct {
	BaseModel
	CreatorID uint32 `gorm:"column:creator;type:int unsigned;not null;comment:创建者" json:"creator_id,omitempty"`
	TeamID    uint32 `gorm:"column:team_id;type:int unsigned;not null;comment:团队ID" json:"team_id,omitempty"`
}

func (u *TeamModel) BeforeCreate(tx *gorm.DB) (err error) {
	var exist bool
	u.TeamID, exist = permission.GetTeamIDByContext(u.GetContext())
	if !exist || u.TeamID == 0 {
		return merr.ErrorInternalServerError("team id not found")
	}
	u.CreatorID, exist = permission.GetUserIDByContext(u.GetContext())
	if !exist || u.CreatorID == 0 {
		return merr.ErrorInternalServerError("user id not found")
	}
	tx.WithContext(u.GetContext())
	return
}

func (u *TeamModel) BeforeUpdate(tx *gorm.DB) (err error) {
	var exist bool
	u.TeamID, exist = permission.GetTeamIDByContext(u.GetContext())
	if !exist || u.TeamID == 0 {
		return merr.ErrorInternalServerError("team id not found")
	}
	tx.WithContext(u.GetContext()).Where(`team_id = ?`, u.TeamID)
	return
}

func (u *TeamModel) BeforeSave(tx *gorm.DB) (err error) {
	var exist bool
	u.TeamID, exist = permission.GetTeamIDByContext(u.GetContext())
	if !exist || u.TeamID == 0 {
		return merr.ErrorInternalServerError("team id not found")
	}
	tx.WithContext(u.GetContext()).Where(`team_id = ?`, u.TeamID)
	return
}

func (u *TeamModel) BeforeDelete(tx *gorm.DB) (err error) {
	var exist bool
	u.TeamID, exist = permission.GetTeamIDByContext(u.GetContext())
	if !exist || u.TeamID == 0 {
		return merr.ErrorInternalServerError("team id not found")
	}
	tx.WithContext(u.GetContext()).Where(`team_id = ?`, u.TeamID)
	return
}
