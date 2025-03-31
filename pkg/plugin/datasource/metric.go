package datasource

import (
	"context"
	"time"
)

type (
	MetricQueryResult struct {
		Metric map[string]string `json:"metric"`
		Value  [2]any            `json:"value"`
		Values [][2]any          `json:"values"`
	}

	MetricQueryData struct {
		ResultType string               `json:"resultType"`
		Result     []*MetricQueryResult `json:"result"`
	}

	MetricQueryResponse struct {
		Status    string           `json:"status"`
		Data      *MetricQueryData `json:"data"`
		ErrorType string           `json:"errorType"`
		Error     string           `json:"error"`
	}

	MetricQueryRequest struct {
		Expr      string
		Duration  time.Duration
		StartTime int64
		EndTime   int64
		Step      uint32
	}

	MetricMetadataItem struct {
		// Name metric name
		Name string `json:"name"`
		// Help metric help
		Help string `json:"help"`
		// Type metric type
		Type string `json:"type"`
		// Labels metric labels
		Labels map[string][]string `json:"labels"`
		// Unit metric unit
		Unit string `json:"unit"`
	}

	MetricMetadata struct {
		Metric    []*MetricMetadataItem `json:"metric"`
		Timestamp int64                 `json:"timestamp"`
	}
)

type Metric interface {
	Query(ctx context.Context, req *MetricQueryRequest) (*MetricQueryResponse, error)

	Metadata(ctx context.Context) (<-chan *MetricMetadata, error)
}
