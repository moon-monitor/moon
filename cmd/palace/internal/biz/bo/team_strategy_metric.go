package bo

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

var _ UpdateTeamStrategyParams = (*SaveTeamStrategyParams)(nil)
var _ CreateTeamStrategyParams = (*SaveTeamStrategyParams)(nil)

type CreateTeamStrategyParams interface {
	GetStrategyGroup() do.StrategyGroup
	GetName() string
	GetRemark() string
	GetStrategyType() vobj.StrategyType
	GetReceiverRoutes() []do.NoticeGroup
	Validate() error
}

type UpdateTeamStrategyParams interface {
	GetStrategy() do.Strategy
	CreateTeamStrategyParams
}

type SaveTeamStrategyParams struct {
	StrategyGroupID uint32
	ID              uint32
	Name            string
	Remark          string
	StrategyType    vobj.StrategyType
	ReceiverRoutes  []uint32

	strategyDo     do.Strategy
	strategyGroup  do.StrategyGroup
	receiverRoutes []do.NoticeGroup
}

// GetStrategyGroup implements TeamStrategyParams.
func (s *SaveTeamStrategyParams) GetStrategyGroup() do.StrategyGroup {
	return s.strategyGroup
}

// GetID implements TeamStrategyParams.
func (s *SaveTeamStrategyParams) GetStrategy() do.Strategy {
	return s.strategyDo
}

// GetName implements TeamStrategyParams.
func (s *SaveTeamStrategyParams) GetName() string {
	return s.Name
}

// GetReceiverRoutes implements TeamStrategyParams.
func (s *SaveTeamStrategyParams) GetReceiverRoutes() []do.NoticeGroup {
	return s.receiverRoutes
}

// GetRemark implements TeamStrategyParams.
func (s *SaveTeamStrategyParams) GetRemark() string {
	return s.Remark
}

// GetStrategyType implements TeamStrategyParams.
func (s *SaveTeamStrategyParams) GetStrategyType() vobj.StrategyType {
	return s.StrategyType
}

func (s *SaveTeamStrategyParams) Validate() error {
	if s.StrategyGroupID <= 0 {
		return merr.ErrorParamsError("strategy group id is required")
	}
	if validate.IsNil(s.strategyGroup) || s.strategyGroup.GetID() != s.StrategyGroupID {
		return merr.ErrorParamsError("strategy group is not found")
	}
	if strings.TrimSpace(s.Name) == "" {
		return merr.ErrorParamsError("name is required")
	}
	if !s.StrategyType.Exist() {
		return merr.ErrorParamsError("strategy type is invalid")
	}
	if utf8.RuneCountInString(s.Remark) > 255 {
		return merr.ErrorParamsError("remark is too long")
	}
	if s.ID > 0 && validate.IsNil(s.strategyDo) {
		return merr.ErrorParamsError("strategy is not found")
	}
	if s.ID > 0 && s.strategyDo.GetID() != s.ID {
		return merr.ErrorParamsError("strategy is not found")
	}
	return nil
}

func (s *SaveTeamStrategyParams) ToUpdateTeamStrategyParams(
	strategyGroup do.StrategyGroup,
	strategyDo do.Strategy,
	receiverRoutes []do.NoticeGroup,
) UpdateTeamStrategyParams {
	s.strategyGroup = strategyGroup
	s.strategyDo = strategyDo
	s.receiverRoutes = receiverRoutes
	return s
}

func (s *SaveTeamStrategyParams) ToCreateTeamStrategyParams(
	strategyGroup do.StrategyGroup,
	receiverRoutes []do.NoticeGroup,
) CreateTeamStrategyParams {
	s.strategyGroup = strategyGroup
	s.receiverRoutes = receiverRoutes
	return s
}

var _ CreateTeamMetricStrategyParams = (*SaveTeamMetricStrategyParams)(nil)
var _ UpdateTeamMetricStrategyParams = (*SaveTeamMetricStrategyParams)(nil)

type CreateTeamMetricStrategyParams interface {
	GetStrategy() do.Strategy
	GetExpr() string
	GetLabels() kv.StringMap
	GetAnnotations() kv.StringMap
	GetDatasource() []do.DatasourceMetric
	Validate() error
}

type UpdateTeamMetricStrategyParams interface {
	CreateTeamMetricStrategyParams
	GetID() uint32
}

type SaveTeamMetricStrategyParams struct {
	StrategyID  uint32
	ID          uint32
	Expr        string
	Labels      kv.StringMap
	Annotations kv.StringMap
	Datasource  []uint32

	strategyDo       do.Strategy
	datasourceDos    []do.DatasourceMetric
	strategyMetricDo do.StrategyMetric
}

// GetAnnotations implements UpdateTeamMetricStrategyParams.
func (s *SaveTeamMetricStrategyParams) GetAnnotations() kv.StringMap {
	return s.Annotations
}

// GetDatasource implements UpdateTeamMetricStrategyParams.
func (s *SaveTeamMetricStrategyParams) GetDatasource() []do.DatasourceMetric {
	return s.datasourceDos
}

// GetExpr implements UpdateTeamMetricStrategyParams.
func (s *SaveTeamMetricStrategyParams) GetExpr() string {
	return s.Expr
}

// GetID implements UpdateTeamMetricStrategyParams.
func (s *SaveTeamMetricStrategyParams) GetID() uint32 {
	return s.ID
}

// GetLabels implements UpdateTeamMetricStrategyParams.
func (s *SaveTeamMetricStrategyParams) GetLabels() kv.StringMap {
	return s.Labels
}

// GetStrategy implements UpdateTeamMetricStrategyParams.
func (s *SaveTeamMetricStrategyParams) GetStrategy() do.Strategy {
	return s.strategyDo
}

func (s *SaveTeamMetricStrategyParams) ToCreateTeamMetricStrategyParams(strategyDo do.Strategy, datasourceDos []do.DatasourceMetric) CreateTeamMetricStrategyParams {
	s.strategyDo = strategyDo
	s.datasourceDos = datasourceDos
	return s
}

func (s *SaveTeamMetricStrategyParams) ToUpdateTeamMetricStrategyParams(
	strategyDo do.Strategy,
	datasourceDos []do.DatasourceMetric,
	strategyMetricDo do.StrategyMetric,
) UpdateTeamMetricStrategyParams {
	s.strategyDo = strategyDo
	s.strategyMetricDo = strategyMetricDo
	s.datasourceDos = datasourceDos
	return s
}

func (s *SaveTeamMetricStrategyParams) Validate() error {
	if s.StrategyID <= 0 {
		return merr.ErrorParamsError("strategy id is required")
	}
	if validate.IsNil(s.strategyDo) {
		return merr.ErrorParamsError("strategy is not found")
	}
	if strings.TrimSpace(s.Expr) == "" {
		return merr.ErrorParamsError("expr is required")
	}
	if len(s.Datasource) == 0 {
		return merr.ErrorParamsError("datasource is required")
	}
	if len(s.Annotations) == 0 {
		return merr.ErrorParamsError("annotations is required")
	}
	if s.ID > 0 && (validate.IsNil(s.strategyMetricDo) || s.strategyMetricDo.GetID() != s.ID) {
		return merr.ErrorParamsError("strategy metric is not found")
	}
	if len(s.Datasource) != len(s.datasourceDos) {
		return merr.ErrorParamsError("datasource is not found")
	}

	return nil
}

type LabelNoticeParams struct {
	ID             uint32
	Key            string
	Value          string
	ReceiverRoutes []uint32
}

type SaveTeamMetricStrategyLevelParams struct {
	ID             uint32
	LevelId        uint32
	LevelName      string
	SampleMode     vobj.SampleMode
	Count          int64
	Condition      vobj.ConditionMetric
	Values         []float64
	ReceiverRoutes []uint32
	LabelNotices   []*LabelNoticeParams
	Duration       time.Duration
	Status         vobj.GlobalStatus
}

type SaveTeamMetricStrategyLevelsParams struct {
	StrategyMetricID uint32
	Levels           []*SaveTeamMetricStrategyLevelParams
}

type UpdateTeamStrategiesStatusParams struct {
	StrategyIds []uint32
	Status      vobj.GlobalStatus
}

func (s *UpdateTeamStrategiesStatusParams) Validate() error {
	if len(s.StrategyIds) == 0 {
		return merr.ErrorParamsError("strategy ids is required")
	}
	if !s.Status.Exist() {
		return merr.ErrorParamsError("status is invalid")
	}
	return nil
}

type OperateTeamStrategyParams struct {
	StrategyId uint32
}

func (s *OperateTeamStrategyParams) Validate() error {
	if s.StrategyId <= 0 {
		return merr.ErrorParamsError("strategy id is required")
	}
	return nil
}

type ListTeamStrategyParams struct {
	*PaginationRequest
	Keyword string
	Status  []vobj.GlobalStatus
}

func (l *ListTeamStrategyParams) Validate() error {
	if l.Keyword != "" && utf8.RuneCountInString(l.Keyword) > 20 {
		return merr.ErrorParamsError("keyword is too long")
	}
	if len(l.Status) > 0 {
		for _, status := range l.Status {
			if !status.Exist() {
				return merr.ErrorParamsError("status is invalid")
			}
		}
	}
	return nil
}

func (l *ListTeamStrategyParams) ToListTeamStrategyReply(items []*team.Strategy) *ListTeamStrategyReply {
	return &ListTeamStrategyReply{
		PaginationReply: l.ToReply(),
		Items:           slices.Map(items, func(item *team.Strategy) do.Strategy { return item }),
	}
}

type ListTeamStrategyReply = ListReply[do.Strategy]

type SubscribeTeamStrategy interface {
	GetStrategy() do.Strategy
	GetNoticeType() vobj.NoticeType
}

type SubscribeTeamStrategyParams struct {
	StrategyId uint32
	NoticeType vobj.NoticeType

	strategyDo do.Strategy
}

func (s *SubscribeTeamStrategyParams) GetStrategy() do.Strategy {
	return s.strategyDo
}

func (s *SubscribeTeamStrategyParams) GetNoticeType() vobj.NoticeType {
	return s.NoticeType
}

func (s *SubscribeTeamStrategyParams) Validate() error {
	if s.StrategyId <= 0 {
		return merr.ErrorParamsError("strategy id is required")
	}
	if validate.IsNil(s.strategyDo) {
		return merr.ErrorParamsError("strategy is not found")
	}
	if !s.NoticeType.Exist() {
		return merr.ErrorParamsError("notice type is invalid")
	}
	return nil
}

func (s *SubscribeTeamStrategyParams) ToSubscribeTeamStrategyParams(strategyDo do.Strategy) SubscribeTeamStrategy {
	s.strategyDo = strategyDo
	return s
}

type SubscribeTeamStrategiesParams struct {
	*PaginationRequest
	StrategyId  uint32
	Subscribers []uint32
	NoticeType  vobj.NoticeType
}

func (s *SubscribeTeamStrategiesParams) Validate() error {
	if s.StrategyId <= 0 {
		return merr.ErrorParamsError("strategy id is required")
	}
	if !s.NoticeType.Exist() {
		return merr.ErrorParamsError("notice type is invalid")
	}
	return nil
}

func (s *SubscribeTeamStrategiesParams) ToSubscribeTeamStrategiesReply(items []*team.StrategySubscriber) *SubscribeTeamStrategiesReply {
	return &SubscribeTeamStrategiesReply{
		PaginationReply: s.ToReply(),
		Items:           slices.Map(items, func(item *team.StrategySubscriber) do.TeamStrategySubscriber { return item }),
	}
}

type SubscribeTeamStrategiesReply = ListReply[do.TeamStrategySubscriber]
