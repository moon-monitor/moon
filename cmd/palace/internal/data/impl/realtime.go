package impl

import (
	"context"
	"time"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/event"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
)

func NewRealtime(data *data.Data) repository.Realtime {
	return &realtimeImpl{
		Data: data,
	}
}

type realtimeImpl struct {
	*data.Data
}

func (r *realtimeImpl) getRealtimeTableName(ctx context.Context, alertStartsAt time.Time) (string, error) {
	teamId, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return "", merr.ErrorPermissionDenied("team id not found")
	}
	eventDB, err := r.GetEventDB(teamId)
	if err != nil {
		return "", err
	}
	tableName, err := event.GetRealtimeTableName(teamId, alertStartsAt, eventDB.GetDB())
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Exists implements repository.Realtime.
func (r *realtimeImpl) Exists(ctx context.Context, alert *bo.GetAlertParams) (bool, error) {
	ctx = permission.WithTeamIDContext(ctx, alert.TeamID)
	tx, teamId, err := getTeamEventQuery(ctx, r)
	if err != nil {
		return false, err
	}
	tableName, err := r.getRealtimeTableName(ctx, alert.StartsAt)
	if err != nil {
		return false, err
	}
	realtimeQuery := tx.Realtime.Table(tableName)
	wrappers := []gen.Condition{
		realtimeQuery.Fingerprint.Eq(alert.Fingerprint),
		realtimeQuery.TeamID.Eq(teamId),
	}

	count, err := realtimeQuery.WithContext(ctx).
		Where(wrappers...).
		Limit(1).
		Count()
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// GetAlert implements repository.Realtime.
func (r *realtimeImpl) GetAlert(ctx context.Context, alert *bo.GetAlertParams) (do.Realtime, error) {
	ctx = permission.WithTeamIDContext(ctx, alert.TeamID)
	tx, teamId, err := getTeamEventQuery(ctx, r)
	if err != nil {
		return nil, err
	}
	tableName, err := r.getRealtimeTableName(ctx, alert.StartsAt)
	if err != nil {
		return nil, err
	}
	realtimeQuery := tx.Realtime.Table(tableName)
	wrappers := []gen.Condition{
		realtimeQuery.Fingerprint.Eq(alert.Fingerprint),
		realtimeQuery.TeamID.Eq(teamId),
	}

	realtimeDo, err := realtimeQuery.WithContext(ctx).
		Where(wrappers...).
		First()
	if err != nil {
		return nil, realtimeNotFound(err)
	}
	return realtimeDo, nil
}

// CreateAlert implements repository.Realtime.
func (r *realtimeImpl) CreateAlert(ctx context.Context, alert *bo.Alert) error {
	ctx = permission.WithTeamIDContext(ctx, alert.TeamID)
	tx, teamId, err := getTeamEventQuery(ctx, r)
	if err != nil {
		return err
	}

	tableName, err := r.getRealtimeTableName(ctx, alert.StartsAt)
	if err != nil {
		return err
	}
	realtimeMutation := tx.Realtime.Table(tableName)
	realtimeDo := &event.Realtime{
		TeamID:       teamId,
		Fingerprint:  alert.Fingerprint,
		Labels:       alert.Labels,
		Summary:      alert.Summary,
		Description:  alert.Description,
		Value:        alert.Value,
		Status:       alert.Status,
		GeneratorURL: alert.GeneratorURL,
		StartsAt:     alert.StartsAt,
		EndsAt:       alert.EndsAt,
	}
	return realtimeMutation.WithContext(ctx).Create(realtimeDo)
}

// UpdateAlert implements repository.Realtime.
func (r *realtimeImpl) UpdateAlert(ctx context.Context, alert *bo.Alert) error {
	ctx = permission.WithTeamIDContext(ctx, alert.TeamID)
	tx, teamId, err := getTeamEventQuery(ctx, r)
	if err != nil {
		return err
	}
	tableName, err := r.getRealtimeTableName(ctx, alert.StartsAt)
	if err != nil {
		return err
	}
	realtimeMutation := tx.Realtime.Table(tableName)
	wrappers := []gen.Condition{
		realtimeMutation.Fingerprint.Eq(alert.Fingerprint),
		realtimeMutation.TeamID.Eq(teamId),
	}
	mutations := []field.AssignExpr{
		realtimeMutation.Status.Value(alert.Status.GetValue()),
		realtimeMutation.GeneratorURL.Value(alert.GeneratorURL),
	}
	if alert.Status.IsResolved() {
		mutations = append(mutations, realtimeMutation.EndsAt.Value(alert.EndsAt))
	}
	_, err = realtimeMutation.WithContext(ctx).
		Where(wrappers...).
		UpdateSimple(mutations...)
	if err != nil {
		return err
	}
	return nil
}
