package bo

type MetricQueryValue struct {
	Value     float64
	Timestamp int64
}

type MetricQueryRangeReply struct {
	Labels     Label
	Values     []*MetricQueryValue
	ResultType string `json:"resultType"`
}

type MetricQueryReply struct {
	Labels     Label
	Value      *MetricQueryValue
	ResultType string `json:"resultType"`
}

type MetricItem struct {
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
