package build

import (
	"github.com/moon-monitor/moon/cmd/laurel/internal/biz/bo"
	apicommon "github.com/moon-monitor/moon/pkg/api/laurel/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToCounterMetricVecs(counterVecs []*apicommon.MetricCounterVec) []*bo.CounterMetricVec {
	return slices.Map(counterVecs, ToCounterMetricVec)
}

func ToGaugeMetricVecs(gaugeVecs []*apicommon.MetricGaugeVec) []*bo.GaugeMetricVec {
	return slices.Map(gaugeVecs, ToGaugeMetricVec)
}

func ToHistogramMetricVecs(histogramVecs []*apicommon.MetricHistogramVec) []*bo.HistogramMetricVec {
	return slices.Map(histogramVecs, ToHistogramMetricVec)
}

func ToSummaryMetricVecs(summaryVecs []*apicommon.MetricSummaryVec) []*bo.SummaryMetricVec {
	return slices.Map(summaryVecs, ToSummaryMetricVec)
}

func ToCounterMetricVec(counterVec *apicommon.MetricCounterVec) *bo.CounterMetricVec {
	if validate.IsNil(counterVec) {
		return nil
	}
	return &bo.CounterMetricVec{
		Namespace: counterVec.GetNamespace(),
		SubSystem: counterVec.GetSubSystem(),
		Name:      counterVec.GetName(),
		Labels:    counterVec.GetLabels(),
		Help:      counterVec.GetHelp(),
	}
}

func ToGaugeMetricVec(gaugeVec *apicommon.MetricGaugeVec) *bo.GaugeMetricVec {
	if validate.IsNil(gaugeVec) {
		return nil
	}
	return &bo.GaugeMetricVec{
		Namespace: gaugeVec.GetNamespace(),
		SubSystem: gaugeVec.GetSubSystem(),
		Name:      gaugeVec.GetName(),
		Labels:    gaugeVec.GetLabels(),
		Help:      gaugeVec.GetHelp(),
	}
}

func ToHistogramMetricVec(histogramVec *apicommon.MetricHistogramVec) *bo.HistogramMetricVec {
	if validate.IsNil(histogramVec) {
		return nil
	}
	return &bo.HistogramMetricVec{
		Namespace:                       histogramVec.GetNamespace(),
		SubSystem:                       histogramVec.GetSubSystem(),
		Name:                            histogramVec.GetName(),
		Labels:                          histogramVec.GetLabels(),
		Help:                            histogramVec.GetHelp(),
		Buckets:                         histogramVec.GetBuckets(),
		NativeHistogramBucketFactor:     histogramVec.GetNativeHistogramBucketFactor(),
		NativeHistogramZeroThreshold:    histogramVec.GetNativeHistogramZeroThreshold(),
		NativeHistogramMaxBucketNumber:  histogramVec.GetNativeHistogramMaxBucketNumber(),
		NativeHistogramMinResetDuration: histogramVec.GetNativeHistogramMinResetDuration(),
		NativeHistogramMaxZeroThreshold: histogramVec.GetNativeHistogramMaxZeroThreshold(),
		NativeHistogramMaxExemplars:     histogramVec.GetNativeHistogramMaxExemplars(),
		NativeHistogramExemplarTTL:      histogramVec.GetNativeHistogramExemplarTTL(),
	}
}

func ToSummaryMetricVec(summaryVec *apicommon.MetricSummaryVec) *bo.SummaryMetricVec {
	if validate.IsNil(summaryVec) {
		return nil
	}
	objectivesList := summaryVec.GetObjectives()
	objectives := make(map[float64]float64, len(objectivesList))
	for _, objective := range objectivesList {
		objectives[objective.GetQuantile()] = objective.GetValue()
	}
	return &bo.SummaryMetricVec{
		Namespace:  summaryVec.GetNamespace(),
		SubSystem:  summaryVec.GetSubSystem(),
		Name:       summaryVec.GetName(),
		Labels:     summaryVec.GetLabels(),
		Help:       summaryVec.GetHelp(),
		Objectives: objectives,
		MaxAge:     summaryVec.GetMaxAge(),
		AgeBuckets: summaryVec.GetAgeBuckets(),
		BufCap:     summaryVec.GetBufCap(),
	}
}
