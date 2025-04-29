package impl

import (
	"context"
	"strings"
	"time"

	"gorm.io/gen"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/timex"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewSendMessageLog(data *data.Data) repository.SendMessageLog {
	return &sendMessageLogImpl{
		Data: data,
	}
}

type sendMessageLogImpl struct {
	*data.Data
}

func (s *sendMessageLogImpl) Retry(ctx context.Context, params *bo.RetrySendMessageParams) error {
	if params.TeamID > 0 {
		return s.retryTeamSendMessageLog(ctx, params)
	}
	return s.retrySystemSendMessageLog(ctx, params)
}

func (s *sendMessageLogImpl) getTeamSendMessageLogTableName(ctx context.Context, sendAt time.Time) (string, error) {
	teamId, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return "", merr.ErrorPermissionDenied("team id not found")
	}
	bizDB, err := s.GetBizDB(teamId)
	if err != nil {
		return "", err
	}
	return team.GetSendMessageLogTableName(bizDB.GetDB(), sendAt, teamId)
}

func (s *sendMessageLogImpl) getSystemSendMessageLogTableName(_ context.Context, sendAt time.Time) (string, error) {
	tx := s.GetMainDB()
	return system.GetSendMessageLogTableName(tx.GetDB(), sendAt)
}

func (s *sendMessageLogImpl) retryTeamSendMessageLog(ctx context.Context, params *bo.RetrySendMessageParams) error {
	tx, teamId, err := getTeamBizQuery(ctx, s)
	if err != nil {
		return err
	}
	tableName, err := s.getTeamSendMessageLogTableName(ctx, params.SendAt)
	if err != nil {
		return err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	wrapper := sendMessageLogTx.WithContext(ctx)
	wrapper = wrapper.Where(sendMessageLogTx.TeamID.Eq(teamId), sendMessageLogTx.RequestID.Eq(params.RequestID))
	_, err = wrapper.UpdateSimple(sendMessageLogTx.RetryCount.Add(1))
	return err
}

func (s *sendMessageLogImpl) retrySystemSendMessageLog(ctx context.Context, params *bo.RetrySendMessageParams) error {
	tx := getMainQuery(ctx, s)
	tableName, err := s.getSystemSendMessageLogTableName(ctx, params.SendAt)
	if err != nil {
		return err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	wrapper := sendMessageLogTx.WithContext(ctx)
	wrapper = wrapper.Where(sendMessageLogTx.RequestID.Eq(params.RequestID))
	_, err = wrapper.UpdateSimple(sendMessageLogTx.RetryCount.Add(1))
	return err
}

// List implements repository.SendMessageLog.
func (s *sendMessageLogImpl) List(ctx context.Context, params *bo.ListSendMessageLogParams) (*bo.ListSendMessageLogReply, error) {
	if params.TeamID > 0 {
		return s.listTeamSendMessageLog(ctx, params)
	}
	return s.listSystemSendMessageLog(ctx, params)
}

// Get implements repository.SendMessageLog.
func (s *sendMessageLogImpl) Get(ctx context.Context, params *bo.GetSendMessageLogParams) (do.SendMessageLog, error) {
	if params.TeamID > 0 {
		return s.getTeamSendMessageLog(ctx, params)
	}
	return s.getSystemSendMessageLog(ctx, params)
}

// UpdateStatus implements repository.SendMessageLog.
func (s *sendMessageLogImpl) UpdateStatus(ctx context.Context, params *bo.UpdateSendMessageLogStatusParams) error {
	if params.TeamID > 0 {
		return s.updateTeamSendMessageLog(ctx, params)
	}
	return s.updateSystemSendMessageLog(ctx, params)
}

func (s *sendMessageLogImpl) Create(ctx context.Context, params *bo.CreateSendMessageLogParams) error {
	if params.TeamID > 0 {
		return s.createTeamSendMessageLog(ctx, params)
	}
	return s.createSystemSendMessageLog(ctx, params)
}

func (s *sendMessageLogImpl) createTeamSendMessageLog(ctx context.Context, params *bo.CreateSendMessageLogParams) error {
	sendMessageLog := &team.SendMessageLog{
		TeamID:      params.TeamID,
		MessageType: params.MessageType,
		Message:     params.Message.String(),
		RequestID:   params.RequestID,
		Status:      vobj.SendMessageStatusSending,
		RetryCount:  0,
		Error:       "",
	}
	sendMessageLog.WithContext(ctx)
	tx, _, err := getTeamBizQuery(ctx, s)
	if err != nil {
		return err
	}
	tableName, err := s.getTeamSendMessageLogTableName(ctx, params.SendAt)
	if err != nil {
		return err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	return sendMessageLogTx.WithContext(ctx).Create(sendMessageLog)
}

func (s *sendMessageLogImpl) createSystemSendMessageLog(ctx context.Context, params *bo.CreateSendMessageLogParams) error {
	sendMessageLog := &system.SendMessageLog{
		MessageType: params.MessageType,
		Message:     params.Message.String(),
		RequestID:   params.RequestID,
		Status:      vobj.SendMessageStatusSending,
	}
	sendMessageLog.WithContext(ctx)
	tx := getMainQuery(ctx, s)
	tableName, err := s.getSystemSendMessageLogTableName(ctx, params.SendAt)
	if err != nil {
		return err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	return sendMessageLogTx.WithContext(ctx).Create(sendMessageLog)
}

func (s *sendMessageLogImpl) getTeamSendMessageLog(ctx context.Context, params *bo.GetSendMessageLogParams) (do.SendMessageLog, error) {
	ctx = permission.WithTeamIDContext(ctx, params.TeamID)
	tx, teamId, err := getTeamBizQuery(ctx, s)
	if err != nil {
		return nil, err
	}
	tableName, err := s.getTeamSendMessageLogTableName(ctx, params.SendAt)
	if err != nil {
		return nil, err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	wrapper := sendMessageLogTx.WithContext(ctx)
	wrappers := []gen.Condition{
		sendMessageLogTx.TeamID.Eq(teamId),
		sendMessageLogTx.RequestID.Eq(params.RequestID),
	}
	sendMessageLog, err := wrapper.Where(wrappers...).First()
	if err != nil {
		return nil, sendMessageLogNotFound(err)
	}
	return sendMessageLog, nil
}

func (s *sendMessageLogImpl) getSystemSendMessageLog(ctx context.Context, params *bo.GetSendMessageLogParams) (do.SendMessageLog, error) {
	tx := getMainQuery(ctx, s)
	tableName, err := s.getSystemSendMessageLogTableName(ctx, params.SendAt)
	if err != nil {
		return nil, err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	wrapper := sendMessageLogTx.WithContext(ctx)
	wrappers := []gen.Condition{
		sendMessageLogTx.RequestID.Eq(params.RequestID),
	}
	sendMessageLog, err := wrapper.Where(wrappers...).First()
	if err != nil {
		return nil, sendMessageLogNotFound(err)
	}
	return sendMessageLog, nil
}

func (s *sendMessageLogImpl) updateTeamSendMessageLog(ctx context.Context, params *bo.UpdateSendMessageLogStatusParams) error {
	ctx = permission.WithTeamIDContext(ctx, params.TeamID)
	tx, teamId, err := getTeamBizQuery(ctx, s)
	if err != nil {
		return err
	}
	tableName, err := s.getTeamSendMessageLogTableName(ctx, params.SendAt)
	if err != nil {
		return err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	wrapper := sendMessageLogTx.WithContext(ctx)
	wrappers := []gen.Condition{
		sendMessageLogTx.TeamID.Eq(teamId),
		sendMessageLogTx.RequestID.Eq(params.RequestID),
	}
	sendMessageLog, err := wrapper.Where(wrappers...).First()
	if err != nil {
		return sendMessageLogNotFound(err)
	}
	sendMessageLog.WithContext(ctx)
	sendMessageLog.Status = params.Status
	sendMessageLog.Error = strings.Join([]string{sendMessageLog.Error, params.Error}, "\n")
	return wrapper.Save(sendMessageLog)
}

func (s *sendMessageLogImpl) updateSystemSendMessageLog(ctx context.Context, params *bo.UpdateSendMessageLogStatusParams) error {
	tx := getMainQuery(ctx, s)
	tableName, err := s.getSystemSendMessageLogTableName(ctx, params.SendAt)
	if err != nil {
		return err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	wrapper := sendMessageLogTx.WithContext(ctx)
	wrappers := []gen.Condition{
		sendMessageLogTx.RequestID.Eq(params.RequestID),
	}
	sendMessageLog, err := wrapper.Where(wrappers...).First()
	if err != nil {
		return sendMessageLogNotFound(err)
	}
	sendMessageLog.WithContext(ctx)
	sendMessageLog.Status = params.Status
	sendMessageLog.Error = strings.Join([]string{sendMessageLog.Error, params.Error}, "\n")
	return wrapper.Save(sendMessageLog)
}

func (s *sendMessageLogImpl) listTeamSendMessageLog(ctx context.Context, params *bo.ListSendMessageLogParams) (*bo.ListSendMessageLogReply, error) {
	tx, teamId, err := getTeamBizQuery(ctx, s)
	if err != nil {
		return nil, err
	}
	sendAt := timex.Now()
	if len(params.TimeRange) == 2 {
		sendAt = params.TimeRange[0]
	}
	tableName, err := s.getTeamSendMessageLogTableName(ctx, sendAt)
	if err != nil {
		return nil, err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	wrapper := sendMessageLogTx.WithContext(ctx).Where(sendMessageLogTx.TeamID.Eq(teamId))
	if validate.TextIsNotNull(params.Keyword) {
		wrapper = wrapper.Where(sendMessageLogTx.Message.Like(params.Keyword))
	}
	if validate.TextIsNotNull(params.RequestID) {
		wrapper = wrapper.Where(sendMessageLogTx.RequestID.Eq(params.RequestID))
	}
	if !params.MessageType.IsUnknown() {
		wrapper = wrapper.Where(sendMessageLogTx.MessageType.Eq(params.MessageType.GetValue()))
	}
	if len(params.TimeRange) == 2 {
		wrapper = wrapper.Where(sendMessageLogTx.CreatedAt.Between(params.TimeRange[0], params.TimeRange[1]))
	}
	if !params.Status.IsUnknown() {
		wrapper = wrapper.Where(sendMessageLogTx.Status.Eq(params.Status.GetValue()))
	}
	if validate.IsNotNil(params.PaginationRequest) {
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		params.WithTotal(total)
		wrapper = wrapper.Limit(int(params.Limit)).Offset(params.Offset())
	}
	sendMessageLogs, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	rows := slices.Map(sendMessageLogs, func(log *team.SendMessageLog) do.SendMessageLog {
		return log
	})
	return params.ToListSendMessageLogReply(rows), nil
}

func (s *sendMessageLogImpl) listSystemSendMessageLog(ctx context.Context, params *bo.ListSendMessageLogParams) (*bo.ListSendMessageLogReply, error) {
	tx := getMainQuery(ctx, s)
	sendAt := timex.Now()
	if len(params.TimeRange) == 2 {
		sendAt = params.TimeRange[0]
	}
	tableName, err := s.getSystemSendMessageLogTableName(ctx, sendAt)
	if err != nil {
		return nil, err
	}
	sendMessageLogTx := tx.SendMessageLog.Table(tableName)
	wrapper := sendMessageLogTx.WithContext(ctx)
	if validate.TextIsNotNull(params.Keyword) {
		wrapper = wrapper.Where(sendMessageLogTx.Message.Like(params.Keyword))
	}
	if validate.TextIsNotNull(params.RequestID) {
		wrapper = wrapper.Where(sendMessageLogTx.RequestID.Eq(params.RequestID))
	}
	if !params.MessageType.IsUnknown() {
		wrapper = wrapper.Where(sendMessageLogTx.MessageType.Eq(params.MessageType.GetValue()))
	}
	if len(params.TimeRange) == 2 {
		wrapper = wrapper.Where(sendMessageLogTx.CreatedAt.Between(params.TimeRange[0], params.TimeRange[1]))
	}
	if !params.Status.IsUnknown() {
		wrapper = wrapper.Where(sendMessageLogTx.Status.Eq(params.Status.GetValue()))
	}
	if validate.IsNotNil(params.PaginationRequest) {
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		params.WithTotal(total)
		wrapper = wrapper.Limit(int(params.Limit)).Offset(params.Offset())
	}
	sendMessageLogs, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	rows := slices.Map(sendMessageLogs, func(log *system.SendMessageLog) do.SendMessageLog {
		return log
	})
	return params.ToListSendMessageLogReply(rows), nil
}
