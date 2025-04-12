package event

import (
	"time"
)

const tableNameRealtime = "realtime"

type Realtime struct {
	ID        uint32    `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"`
}

func (r *Realtime) TableName() string {
	return tableNameRealtime
}
