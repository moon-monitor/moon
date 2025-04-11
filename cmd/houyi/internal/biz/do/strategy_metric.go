package do

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

type MetricRule struct {
	TeamId        uint32
	Datasource    string
	StrategyId    uint32
	LevelId       uint32
	Receiver      []string
	LabelReceiver []*LabelNotices
	Expr          string
	Labels        *label.Label
	Annotations   *label.Annotation
	Duration      time.Duration
	Count         int64
	Values        []float64
	SampleMode    common.MetricStrategyItem_SampleMode
	Condition     common.MetricStrategyItem_Condition
	Enable        bool `json:"enable"`
}

func (m *MetricRule) GetEnable() bool {
	if m == nil {
		return false
	}
	return m.Enable
}

func (m *MetricRule) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *MetricRule) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *MetricRule) UniqueKey() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("%d:%d:%s", m.TeamId, m.StrategyId, m.Datasource)
}

func (m *MetricRule) GetTeamId() uint32 {
	if m == nil {
		return 0
	}
	return m.TeamId
}

func (m *MetricRule) GetDatasource() string {
	if m == nil {
		return ""
	}
	return m.Datasource
}

func (m *MetricRule) GetStrategyId() uint32 {
	if m == nil {
		return 0
	}
	return m.StrategyId
}

func (m *MetricRule) GetLevelId() uint32 {
	if m == nil {
		return 0
	}
	return m.LevelId
}

func (m *MetricRule) GetReceiverRoutes() []string {
	if m == nil {
		return nil
	}
	return m.Receiver
}

func (m *MetricRule) GetLabelReceiverRoutes() []bo.LabelNotices {
	if m == nil {
		return nil
	}
	return slices.Map(m.LabelReceiver, func(v *LabelNotices) bo.LabelNotices {
		return v
	})
}

func (m *MetricRule) GetExpr() string {
	if m == nil {
		return ""
	}
	return m.Expr
}

func (m *MetricRule) GetLabels() *label.Label {
	if m == nil {
		return nil
	}
	return m.Labels
}

func (m *MetricRule) GetAnnotations() *label.Annotation {
	if m == nil {
		return nil
	}
	return m.Annotations
}

func (m *MetricRule) GetDuration() time.Duration {
	if m == nil {
		return 0
	}
	return m.Duration
}

func (m *MetricRule) GetCount() int64 {
	if m == nil {
		return 0
	}
	return m.Count
}

func (m *MetricRule) GetValues() []float64 {
	if m == nil {
		return nil
	}
	return m.Values
}

func (m *MetricRule) GetSampleMode() common.MetricStrategyItem_SampleMode {
	if m == nil {
		return common.MetricStrategyItem_SampleMode_For
	}
	return m.SampleMode
}

func (m *MetricRule) GetCondition() common.MetricStrategyItem_Condition {
	if m == nil {
		return common.MetricStrategyItem_Condition_EQ
	}
	return m.Condition
}

func (m *MetricRule) GetExt() kv.Map[string, any] {
	if m == nil {
		return nil
	}
	return map[string]any{}
}
