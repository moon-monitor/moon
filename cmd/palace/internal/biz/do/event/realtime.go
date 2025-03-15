package event

import (
	"time"
)

const tableNameRealtime = "realtime"

type Realtime struct {
	ID        uint32    `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at,omitempty"`
}

func (r *Realtime) TableName() string {
	return tableNameRealtime
}
