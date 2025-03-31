package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/status"

	"github.com/moon-monitor/moon/pkg/plugin/datasource"
	"github.com/moon-monitor/moon/pkg/util/httpx"
)

var _ datasource.Metric = (*Prometheus)(nil)

const (
	// prometheusAPIV1Query 查询接口
	prometheusAPIV1Query = "/api/v1/query"
	// prometheusAPIV1QueryRange 查询接口
	prometheusAPIV1QueryRange = "/api/v1/query_range"
	// prometheusAPIV1Metadata 元数据查询接口
	prometheusAPIV1Metadata = "/api/v1/metadata"
	// prometheusAPIV1Series /api/v1/series
	prometheusAPIV1Series = "/api/v1/series"
)

type Config interface {
	GetEndpoint() string
	GetBasicAuth() datasource.BasicAuth
}

func New(c Config, logger log.Logger) *Prometheus {
	return &Prometheus{
		c:      c,
		helper: log.NewHelper(log.With(logger, "module", "plugin.datasource.prometheus")),
	}
}

type Prometheus struct {
	c      Config
	helper *log.Helper
}

func (p *Prometheus) Query(ctx context.Context, req *datasource.PromQueryRequest) ([]*datasource.PromQueryResponse, error) {
	if req.StartTime > 0 && req.EndTime > 0 {
		return p.queryRange(ctx, req.Expr, req.StartTime, req.EndTime, req.Step)
	}
	return p.query(ctx, req.Expr, int64(req.Duration))
}

func (p *Prometheus) Metadata(ctx context.Context) (<-chan *datasource.MetricMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Prometheus) query(ctx context.Context, expr string, duration int64) ([]*datasource.PromQueryResponse, error) {
	params := httpx.ParseQuery(map[string]any{
		"query": expr,
		"time":  duration,
	})

	hx := httpx.NewClient().WithContext(ctx)
	hx = hx.WithHeader(http.Header{
		"Accept":          []string{"*/*"},
		"Accept-Language": []string{"zh-CN,zh;q=0.9"},
	})
	if p.c.GetBasicAuth() != nil {
		basicAuth := p.c.GetBasicAuth()
		hx = hx.WithBasicAuth(basicAuth.GetUsername(), basicAuth.GetPassword())
	}
	api, err := url.JoinPath(p.c.GetEndpoint(), prometheusAPIV1Query)
	if err != nil {
		return nil, err
	}
	getResponse, err := hx.Get(fmt.Sprintf("%s?%s", api, params))
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()
	var allResp datasource.PromQueryResponse
	if err = json.NewDecoder(getResponse.Body).Decode(&allResp); err != nil {
		return nil, err
	}
	return &allResp, nil

	if allResp.Error != "" {
		return nil, status.Errorf(400, "query error: %s", allResp.Error)
	}
	data := allResp.Data
	if data == nil {
		return []*datasource.PromQueryResponse(nil), nil
	}
	result := make([]*datasource.PromQueryData, 0, len(data.Result))
	for _, v := range data.Result {
		value := v.Value
		ts, tsAssertOk := strconv.ParseFloat(fmt.Sprintf("%v", value[0]), 64)
		if tsAssertOk != nil {
			continue
		}
		metricValue, parseErr := strconv.ParseFloat(fmt.Sprintf("%v", value[1]), 64)
		if parseErr != nil {
			continue
		}
		result = append(result, &datasource.PromQueryData{
			Labels: v.Metric,
			Value: &QueryValue{
				Value:     metricValue,
				Timestamp: int64(ts),
			},
			ResultType: data.ResultType,
		})
	}

	return result, nil
}

func (p *Prometheus) queryRange(ctx context.Context, expr string, start, end int64, step uint32) ([]*datasource.PromQueryResponse, error) {
	st := step
	if step == 0 {
		st = p.step
	}
	params := httpx.ParseQuery(map[string]any{
		"query": expr,
		"start": start,
		"end":   end,
		"step":  st,
	})

	hx := httpx.NewHTTPX()
	hx.SetHeader(map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
	})
	if p.basicAuth != nil {
		hx = hx.SetBasicAuth(p.basicAuth.Username, p.basicAuth.Password)
	}
	api, err := url.JoinPath(p.endpoint, p.prometheusAPIV1QueryRange)
	if err != nil {
		return nil, err
	}
	getResponse, err := hx.GETWithContext(ctx, fmt.Sprintf("%s?%s", api, params))
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()
	var allResp PromQueryResponse
	if err = types.NewDecoder(getResponse.Body).Decode(&allResp); err != nil {
		return nil, err
	}
	if allResp.Error != "" {
		return nil, status.Errorf(400, "query error: %s", allResp.Error)
	}
	data := allResp.Data
	if types.IsNil(data) {
		return []*QueryResponse(nil), nil
	}
	result := make([]*QueryResponse, 0, len(data.Result))
	for _, v := range data.Result {
		values := make([]*QueryValue, 0, len(v.Values))
		for _, vv := range v.Values {
			ts, tsAssertOk := strconv.ParseFloat(fmt.Sprintf("%v", vv[0]), 64)
			if tsAssertOk != nil {
				continue
			}
			metricValue, parseErr := strconv.ParseFloat(fmt.Sprintf("%v", vv[1]), 64)
			if parseErr != nil {
				continue
			}
			values = append(values, &QueryValue{
				Value:     metricValue,
				Timestamp: int64(ts),
			})
		}

		result = append(result, &QueryResponse{
			Labels:     v.Metric,
			Values:     values,
			ResultType: data.ResultType,
		})
	}
	return result, nil
}
