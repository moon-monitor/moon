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

type Base interface {
	GetID() uint32
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() soft_delete.DeletedAt
	GetContext() context.Context
	WithContext(context.Context)
}

type Creator interface {
	Base
	GetCreatorID() uint32
	GetCreator() User
	WithCreator(func(ctx context.Context, creatorID uint32) (User, error)) error
}

type TeamBase interface {
	Creator
	GetTeamID() uint32
}

var _ Base = (*BaseModel)(nil)

// BaseModel gorm base model
type BaseModel struct {
	ctx context.Context `gorm:"-"`

	ID        uint32                `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time             `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time             `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at,omitempty"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint;not null;default:0;" json:"deleted_at,omitempty"`
}

func (u *BaseModel) GetID() uint32 {
	if u == nil {
		return 0
	}
	return u.ID
}

func (u *BaseModel) GetCreatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.CreatedAt
}

func (u *BaseModel) GetUpdatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.UpdatedAt
}

func (u *BaseModel) GetDeletedAt() soft_delete.DeletedAt {
	if u == nil {
		return 0
	}
	return u.DeletedAt
}

// WithContext set context
func (u *BaseModel) WithContext(ctx context.Context) {
	u.ctx = ctx
}

// GetContext get context
func (u *BaseModel) GetContext() context.Context {
	if u.ctx == nil {
		panic("context is nil")
	}
	return u.ctx
}

var _ Creator = (*CreatorModel)(nil)

type CreatorModel struct {
	BaseModel
	CreatorID uint32 `gorm:"column:creator;type:int unsigned;not null;comment:创建者" json:"creator_id,omitempty"`

	creator User
}

func (u *CreatorModel) GetCreatorID() uint32 {
	if u == nil {
		return 0
	}
	return u.CreatorID
}

func (u *CreatorModel) GetCreator() User {
	if u == nil {
		return nil
	}
	return u.creator
}

func (u *CreatorModel) WithCreator(f func(ctx context.Context, creatorID uint32) (User, error)) error {
	if u == nil || f == nil {
		return nil
	}

	creator, err := f(u.GetContext(), u.CreatorID)
	if err != nil {
		return err
	}

	u.creator = creator
	return nil
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

var _ TeamBase = (*TeamModel)(nil)

type TeamModel struct {
	CreatorModel
	TeamID uint32 `gorm:"column:team_id;type:int unsigned;not null;comment:团队ID" json:"team_id,omitempty"`
}

func (u *TeamModel) GetTeamID() uint32 {
	if u == nil {
		return 0
	}
	return u.TeamID
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
