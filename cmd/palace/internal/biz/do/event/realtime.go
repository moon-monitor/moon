package event

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/kv"
)

var _ do.Realtime = (*Realtime)(nil)

const tableNameRealtime = "team_realtime_alerts"

type Realtime struct {
	ID           uint32           `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt    time.Time        `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt"`
	UpdatedAt    time.Time        `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"`
	TeamID       uint32           `gorm:"column:team_id;type:int;not null;comment:团队ID;uniqueIndex:uk__team_id__fingerprint" json:"teamId"`
	Status       vobj.AlertStatus `gorm:"column:status;type:tinyint;not null;comment:状态" json:"status"`
	Fingerprint  string           `gorm:"column:fingerprint;type:varchar(255);not null;comment:指纹;uniqueIndex:uk__team_id__fingerprint" json:"fingerprint"`
	Labels       kv.StringMap     `gorm:"column:labels;type:text;not null;comment:标签" json:"labels"`
	Summary      string           `gorm:"column:summary;type:text;not null;comment:摘要" json:"summary"`
	Description  string           `gorm:"column:description;type:text;not null;comment:描述" json:"description"`
	Value        string           `gorm:"column:value;type:text;not null;comment:值" json:"value"`
	GeneratorURL string           `gorm:"column:generator_url;type:text;not null;comment:生成URL" json:"generatorURL"`
	StartsAt     time.Time        `gorm:"column:starts_at;type:datetime;not null;default:'0001-01-01 00:00:00';comment:开始时间" json:"startsAt"`
	EndsAt       time.Time        `gorm:"column:ends_at;type:datetime;not null;default:'0001-01-01 00:00:00';comment:结束时间" json:"endsAt"`
}

// GetCreatedAt implements do.Realtime.
func (r *Realtime) GetCreatedAt() time.Time {
	return r.CreatedAt
}

// GetDescription implements do.Realtime.
func (r *Realtime) GetDescription() string {
	return r.Description
}

// GetEndsAt implements do.Realtime.
func (r *Realtime) GetEndsAt() time.Time {
	return r.EndsAt
}

// GetFingerprint implements do.Realtime.
func (r *Realtime) GetFingerprint() string {
	return r.Fingerprint
}

// GetGeneratorURL implements do.Realtime.
func (r *Realtime) GetGeneratorURL() string {
	return r.GeneratorURL
}

// GetID implements do.Realtime.
func (r *Realtime) GetID() uint32 {
	return r.ID
}

// GetLabels implements do.Realtime.
func (r *Realtime) GetLabels() kv.StringMap {
	return r.Labels
}

// GetStartsAt implements do.Realtime.
func (r *Realtime) GetStartsAt() time.Time {
	return r.StartsAt
}

// GetStatus implements do.Realtime.
func (r *Realtime) GetStatus() vobj.AlertStatus {
	return r.Status
}

// GetSummary implements do.Realtime.
func (r *Realtime) GetSummary() string {
	return r.Summary
}

// GetTeamID implements do.Realtime.
func (r *Realtime) GetTeamID() uint32 {
	return r.TeamID
}

// GetUpdatedAt implements do.Realtime.
func (r *Realtime) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}

// GetValue implements do.Realtime.
func (r *Realtime) GetValue() string {
	return r.Value
}

func (r *Realtime) TableName() string {
	return genRealtimeTableName(r.StartsAt, r.TeamID)
}

func createRealtimeTable(tx *gorm.DB, t time.Time, teamId uint32) (err error) {
	tableName := genRealtimeTableName(t, teamId)
	if do.HasTable(tx, tableName) {
		return
	}
	r := &Realtime{
		TeamID:   teamId,
		StartsAt: t,
	}

	if err := do.CreateTable(tx, tableName, r); err != nil {
		return err
	}
	return
}

func genRealtimeTableName(t time.Time, teamId uint32) string {
	offset := time.Monday - t.Weekday()
	weekStart := t.AddDate(0, 0, int(offset))

	return fmt.Sprintf("%s_%d_%s", tableNameRealtime, teamId, weekStart.Format("20060102"))
}

func GetRealtimeTableName(tx *gorm.DB, t time.Time, teamId uint32) (string, error) {
	tableName := genRealtimeTableName(t, teamId)
	if !do.HasTable(tx, tableName) {
		return tableName, createRealtimeTable(tx, t, teamId)
	}
	return tableName, nil
}
