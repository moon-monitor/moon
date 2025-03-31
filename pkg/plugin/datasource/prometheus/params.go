package prometheus

type (
	PromMetricSeriesResponse struct {
		Status    string              `json:"status"`
		Data      []map[string]string `json:"data"`
		Error     string              `json:"error"`
		ErrorType string              `json:"errorType"`
	}

	PromMetricInfo struct {
		Type string `json:"type"`
		Help string `json:"help"`
		Unit string `json:"unit"`
	}

	PromMetadataResponse struct {
		Status string                      `json:"status"`
		Data   map[string][]PromMetricInfo `json:"data"`
	}
)
