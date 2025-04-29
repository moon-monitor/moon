package system

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

var _ do.SendMessageLog = (*SendMessageLog)(nil)

const tableNameSendMessageLog = "sys_send_message_logs"

type SendMessageLog struct {
	do.BaseModel
	SendedAt    time.Time              `gorm:"column:sended_at;type:datetime;not null;comment:发送时间" json:"sended_at,omitempty"`
	MessageType vobj.MessageType       `gorm:"column:message_type;type:tinyint(2);not null;comment:消息类型" json:"message_type,omitempty"`
	Message     string                 `gorm:"column:message;type:text;not null;comment:消息内容" json:"message,omitempty"`
	RequestID   string                 `gorm:"column:request_id;type:varchar(64);not null;comment:请求ID;uniqueIndex:idx__request_id" json:"request_id,omitempty"`
	Status      vobj.SendMessageStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status,omitempty"`
	RetryCount  int                    `gorm:"column:retry_count;type:int unsigned;not null;comment:重试次数" json:"retry_count,omitempty"`
	Error       string                 `gorm:"column:error;type:text;not null;comment:错误信息" json:"error,omitempty"`
}

// GetError implements do.SendMessageLog.
func (s *SendMessageLog) GetError() string {
	if s == nil {
		return ""
	}
	return s.Error
}

// GetMessage implements do.SendMessageLog.
func (s *SendMessageLog) GetMessage() string {
	if s == nil {
		return ""
	}
	return s.Message
}

// GetMessageType implements do.SendMessageLog.
func (s *SendMessageLog) GetMessageType() vobj.MessageType {
	if s == nil {
		return vobj.MessageTypeUnknown
	}
	return s.MessageType
}

// GetRequestID implements do.SendMessageLog.
func (s *SendMessageLog) GetRequestID() string {
	if s == nil {
		return ""
	}
	return s.RequestID
}

// GetRetryCount implements do.SendMessageLog.
func (s *SendMessageLog) GetRetryCount() int32 {
	if s == nil {
		return 0
	}
	return int32(s.RetryCount)
}

// GetStatus implements do.SendMessageLog.
func (s *SendMessageLog) GetStatus() vobj.SendMessageStatus {
	if s == nil {
		return vobj.SendMessageStatusUnknown
	}
	return s.Status
}

// GetTeamID implements do.SendMessageLog.
func (s *SendMessageLog) GetTeamID() uint32 {
	return 0
}

func (s *SendMessageLog) TableName() string {
	return genSendMessageLogTableName(s.SendedAt)
}

func createSendMessageLogTable(tx *gorm.DB, t time.Time) (err error) {
	tableName := genSendMessageLogTableName(t)
	if do.HasTable(tx, tableName) {
		return
	}
	s := &SendMessageLog{SendedAt: t}
	if err := do.CreateTable(tx, tableName, s); err != nil {
		return err
	}
	return
}

func genSendMessageLogTableName(t time.Time) string {
	offset := time.Monday - t.Weekday()
	weekStart := t.AddDate(0, 0, int(offset))

	return fmt.Sprintf("%s_%s", tableNameSendMessageLog, weekStart.Format("20060102"))
}

func GetSendMessageLogTableName(tx *gorm.DB, t time.Time) (string, error) {
	tableName := genSendMessageLogTableName(t)
	if !do.HasTable(tx, tableName) {
		return tableName, createSendMessageLogTable(tx, t)
	}
	return tableName, nil
}
