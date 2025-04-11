package impl

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/event"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/houyi/internal/conf"
	"github.com/moon-monitor/moon/cmd/houyi/internal/data"
	"github.com/moon-monitor/moon/cmd/houyi/internal/data/impl/judge"
	"github.com/moon-monitor/moon/cmd/houyi/internal/data/impl/judge/condition"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/util/cnst"
	"github.com/moon-monitor/moon/pkg/util/pointer"
	"github.com/moon-monitor/moon/pkg/util/template"
)

func NewJudgeRepo(bc *conf.Bootstrap, data *data.Data, logger log.Logger) repository.Judge {
	return &judgeImpl{
		Data:             data,
		evaluateInterval: bc.GetEvaluate().GetInterval().AsDuration() * 2,
		helper:           log.NewHelper(log.With(logger, "module", "data.repo.judge")),
	}
}

type judgeImpl struct {
	*data.Data
	evaluateInterval time.Duration
	helper           *log.Helper
}

func (j *judgeImpl) Metric(_ context.Context, data []bo.MetricJudgeData, rule bo.MetricJudgeRule) ([]bo.Alert, error) {
	conditionType := condition.NewMetricCondition(rule.GetCondition())
	opts := []judge.MetricJudgeOption{
		judge.WithMetricJudgeCondition(conditionType),
		judge.WithMetricJudgeConditionValues(rule.GetValues()),
		judge.WithMetricJudgeConditionCount(rule.GetCount()),
	}
	judgeInstance := judge.NewMetricJudge(rule.GetSampleMode(), opts...)
	alerts := make([]bo.Alert, 0, len(data))
	for _, datum := range data {
		value, ok := judgeInstance.Judge(datum.GetValues())
		if !ok {
			continue
		}
		alert := j.generateAlert(rule, value, datum.GetLabels())
		alerts = append(alerts, alert)
	}
	return alerts, nil
}

func (j *judgeImpl) generateAlert(rule bo.MetricJudgeRule, value bo.MetricJudgeDataValue, originLabels map[string]string) *event.AlertJob {
	ext := rule.GetExt()
	ext.Set(cnst.ExtKeyValues, value.GetValue())
	ext.Set(cnst.ExtKeyTimestamp, value.GetTimestamp())

	labels := rule.GetLabels()
	labelsMap := labels.ToMap()
	for k, v := range labelsMap {
		labelsMap[k] = template.TextFormatterX(v, ext)
	}
	labels = labels.Appends(labelsMap).Appends(originLabels)
	ext.Set(cnst.ExtKeyLabels, labels.ToMap())

	annotations := rule.GetAnnotations()
	summary := template.TextFormatterX(annotations.GetSummary(), ext)
	description := template.TextFormatterX(annotations.GetDescription(), ext)
	annotations.SetSummary(summary)
	annotations.SetDescription(description)

	now := time.Now()
	alert := &event.AlertJob{
		Status:       common.EventStatus_pending,
		Labels:       labels,
		Annotations:  annotations,
		StartsAt:     pointer.Of(time.Unix(value.GetTimestamp(), 0).UTC()),
		EndsAt:       nil,
		GeneratorURL: "",
		Fingerprint:  "",
		Value:        value.GetValue(),
		LastUpdated:  now,
		Duration:     j.evaluateInterval,
	}
	return alert
}
