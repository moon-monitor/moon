package bo

import (
	"time"

	"gorm.io/plugin/soft_delete"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type SaveDictReq struct {
	DictID uint32            `json:"dictID"`
	Key    string            `json:"key"`
	Value  string            `json:"value"`
	Status vobj.GlobalStatus `json:"status"`
	Type   vobj.DictType     `json:"type"`
	Color  string            `json:"color"`
	Lang   string            `json:"lang"`
}

type updateDictReq struct {
	dictItem Dict

	Key    string            `json:"key"`
	Value  string            `json:"value"`
	Status vobj.GlobalStatus `json:"status"`
	Type   vobj.DictType     `json:"type"`
	Color  string            `json:"color"`
	Lang   string            `json:"lang"`
}

func (u *updateDictReq) GetTeamID() uint32 {
	if u == nil || u.dictItem == nil {
		return 0
	}
	return u.dictItem.GetTeamID()
}

func (u *updateDictReq) GetID() uint32 {
	if u == nil || u.dictItem == nil {
		return 0
	}
	return u.dictItem.GetID()
}

func (u *updateDictReq) GetKey() string {
	return u.Key
}

func (u *updateDictReq) GetValue() string {
	return u.Value
}

func (u *updateDictReq) GetStatus() vobj.GlobalStatus {
	return u.Status
}

func (u *updateDictReq) GetType() vobj.DictType {
	return u.Type
}

func (u *updateDictReq) GetColor() string {
	return u.Color
}

func (u *updateDictReq) GetLang() string {
	return u.Lang
}

func (u *updateDictReq) GetCreatedAt() time.Time {
	if u == nil || u.dictItem == nil {
		return time.Now()
	}
	return u.dictItem.GetCreatedAt()
}

func (u *updateDictReq) GetUpdatedAt() time.Time {
	if u == nil || u.dictItem == nil {
		return time.Now()
	}
	return u.dictItem.GetUpdatedAt()
}

func (u *updateDictReq) GetDeletedAt() soft_delete.DeletedAt {
	if u == nil || u.dictItem == nil {
		return 0
	}
	return u.dictItem.GetDeletedAt()
}

func (d *SaveDictReq) WithCreateParams() Dict {
	return &updateDictReq{
		dictItem: nil,
		Key:      d.Key,
		Value:    d.Value,
		Status:   d.Status,
		Type:     d.Type,
		Color:    d.Color,
		Lang:     d.Lang,
	}
}

func (d *SaveDictReq) WithUpdateParams(dictItem Dict) Dict {
	return &updateDictReq{
		dictItem: dictItem,
		Key:      d.Key,
		Value:    d.Value,
		Status:   d.Status,
		Type:     d.Type,
		Color:    d.Color,
		Lang:     d.Lang,
	}
}

type UpdateDictStatusReq struct {
	DictIds []uint32          `json:"dictIds"`
	Status  vobj.GlobalStatus `json:"status"`
}

type OperateOneDictReq struct {
	DictID uint32 `json:"dictID"`
}

type ListDictReq struct {
	*PaginationRequest
	DictTypes []vobj.DictType   `json:"dictTypes"`
	Status    vobj.GlobalStatus `json:"status"`
	Keyword   string            `json:"keyword"`
	Langs     []string          `json:"langs"`
}

type ListDictReply struct {
	Items []Dict
	*PaginationReply
}

func (l *ListDictReq) NewListDictReply(dictItems []Dict) *ListDictReply {
	return &ListDictReply{
		Items:           dictItems,
		PaginationReply: l.ToReply(),
	}
}

type Dict interface {
	GetTeamID() uint32
	GetID() uint32
	GetKey() string
	GetValue() string
	GetStatus() vobj.GlobalStatus
	GetType() vobj.DictType
	GetColor() string
	GetLang() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() soft_delete.DeletedAt
}
