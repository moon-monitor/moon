package prometheus

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/merr"
	"golang.org/x/sync/errgroup"

	"github.com/moon-monitor/moon/pkg/plugin/datasource"
	"github.com/moon-monitor/moon/pkg/util/httpx"
)

var _ datasource.Metric = (*Prometheus)(nil)

const (
	// prometheusAPIV1Query query api v1
	prometheusAPIV1Query = "/api/v1/query"
	// prometheusAPIV1QueryRange query range api v1
	prometheusAPIV1QueryRange = "/api/v1/query_range"
	// prometheusAPIV1Metadata metadata api
	prometheusAPIV1Metadata = "/api/v1/metadata"
	// prometheusAPIV1Series series api
	prometheusAPIV1Series = "/api/v1/series"
)

type Config interface {
	GetEndpoint() string
	GetHeaders() map[string]string
	GetMethod() common.DatasourceQueryMethod
	GetBasicAuth() datasource.BasicAuth
	GetTLS() datasource.TLS
	GetCA() string
	GetScrapeInterval() time.Duration
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

func (p *Prometheus) GetScrapeInterval() time.Duration {
	if p.c.GetScrapeInterval() > 0 {
		return p.c.GetScrapeInterval()
	}
	return 15 * time.Second
}

func (p *Prometheus) Query(ctx context.Context, req *datasource.MetricQueryRequest) (*datasource.MetricQueryResponse, error) {
	if req.StartTime > 0 && req.EndTime > 0 {
		return p.queryRange(ctx, req.Expr, req.StartTime, req.EndTime, req.Step)
	}
	return p.query(ctx, req.Expr, req.Time)
}

func (p *Prometheus) Metadata(ctx context.Context) (<-chan *datasource.MetricMetadata, error) {
	metadataInfo, err := p.metadata(ctx)
	if err != nil {
		return nil, err
	}

	send := make(chan *datasource.MetricMetadata, 20)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				p.helper.Errorw("method", "metadata", "panic", err)
			}
		}()
		defer close(send)
		p.sendMetadata(send, metadataInfo)
	}()

	return send, nil
}

func (p *Prometheus) sendMetadata(send chan<- *datasource.MetricMetadata, metrics map[string][]PromMetricInfo) {
	metricNameMap := make(map[string]PromMetricInfo)
	metricNames := make([]string, 0, len(metrics))
	for metricName := range metrics {
		metricNames = append(metricNames, metricName)
		if len(metrics[metricName]) == 0 {
			continue
		}
		metricNameMap[metricName] = metrics[metricName][0]
	}

	now := time.Now()
	batchNum := 20
	namesLen := len(metricNames)
	eg := new(errgroup.Group)
	for i := 0; i < namesLen; i += batchNum {
		left := i
		right := left + batchNum
		if right > namesLen {
			right = namesLen
		}
		eg.Go(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			seriesInfo, seriesErr := p.series(ctx, now, metricNames[left:right]...)
			if seriesErr != nil {
				log.Warnw("series error", seriesErr)
				return seriesErr
			}

			metricsTmp := make([]*datasource.MetricMetadataItem, 0, right-left)
			for _, metricName := range metricNames[left:right] {
				metricInfo := metricNameMap[metricName]
				item := &datasource.MetricMetadataItem{
					Type:   metricInfo.Type,
					Name:   metricName,
					Help:   metricInfo.Help,
					Unit:   metricInfo.Unit,
					Labels: seriesInfo[metricName],
				}
				metricsTmp = append(metricsTmp, item)
			}
			send <- &datasource.MetricMetadata{
				Metric:    metricsTmp,
				Timestamp: time.Now().Unix(),
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		p.helper.Errorw("method", "metadata", "err", err)
	}
}

func (p *Prometheus) query(ctx context.Context, expr string, t int64) (*datasource.MetricQueryResponse, error) {
	params := httpx.ParseQuery(map[string]any{
		"query": expr,
		"time":  t,
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
	getResponse, err := hx.Get(api, params)
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()
	if getResponse.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(getResponse.Body)
		return nil, merr.ErrorBadRequest("status code: %d => %s", getResponse.StatusCode, string(body))
	}
	var allResp datasource.MetricQueryResponse
	if err = json.NewDecoder(getResponse.Body).Decode(&allResp); err != nil {
		return nil, err
	}
	return &allResp, nil
}

func (p *Prometheus) queryRange(ctx context.Context, expr string, start, end int64, step uint32) (*datasource.MetricQueryResponse, error) {
	params := httpx.ParseQuery(map[string]any{
		"query": expr,
		"start": start,
		"end":   end,
		"step":  step,
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
	api, err := url.JoinPath(p.c.GetEndpoint(), prometheusAPIV1QueryRange)
	if err != nil {
		return nil, err
	}
	getResponse, err := hx.Get(api, params)
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()
	if getResponse.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(getResponse.Body)
		return nil, merr.ErrorBadRequest("status code: %d => %s", getResponse.StatusCode, string(body))
	}
	var allResp datasource.MetricQueryResponse
	if err = json.NewDecoder(getResponse.Body).Decode(&allResp); err != nil {
		return nil, err
	}

	return &allResp, nil
}

func (p *Prometheus) series(ctx context.Context, now time.Time, metricNames ...string) (map[string]map[string][]string, error) {
	start := now.Add(-time.Hour * 24).Format("2006-01-02T15:04:05.000Z")
	end := now.Format("2006-01-02T15:04:05.000Z")

	params := httpx.ParseQuery(map[string]any{
		"start": start,
		"end":   end,
	})

	for _, metricName := range metricNames {
		params.Set("match[]", metricName)
	}

	hx := httpx.NewClient().WithContext(ctx)
	hx = hx.WithHeader(http.Header{
		"Accept":          []string{"*/*"},
		"Accept-Language": []string{"zh-CN,zh;q=0.9"},
	})
	if p.c.GetBasicAuth() != nil {
		basicAuth := p.c.GetBasicAuth()
		hx = hx.WithBasicAuth(basicAuth.GetUsername(), basicAuth.GetPassword())
	}

	api, err := url.JoinPath(p.c.GetEndpoint(), prometheusAPIV1Series)
	if err != nil {
		return nil, err
	}
	getResponse, err := hx.Get(api, params)
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()
	if getResponse.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(getResponse.Body)
		return nil, merr.ErrorBadRequest("status code: %d => %s", getResponse.StatusCode, string(body))
	}
	var allResp PromMetricSeriesResponse
	if err = json.NewDecoder(getResponse.Body).Decode(&allResp); err != nil {
		return nil, err
	}

	res := make(map[string]map[string][]string)
	for _, v := range allResp.Data {
		metricName := v["__name__"]
		if metricName == "" {
			continue
		}
		if _, ok := res[metricName]; !ok {
			res[metricName] = make(map[string][]string)
		}
		for k, val := range v {
			if k == "__name__" {
				continue
			}
			if _, ok := res[metricName][k]; !ok {
				res[metricName][k] = make([]string, 0)
			}
			res[metricName][k] = append(res[metricName][k], val)
		}
	}

	return res, nil
}

func (p *Prometheus) metadata(ctx context.Context) (map[string][]PromMetricInfo, error) {
	hx := httpx.NewClient().WithContext(ctx)
	hx = hx.WithHeader(http.Header{
		"Accept":          []string{"*/*"},
		"Accept-Language": []string{"zh-CN,zh;q=0.9"},
	})
	if p.c.GetBasicAuth() != nil {
		basicAuth := p.c.GetBasicAuth()
		hx = hx.WithBasicAuth(basicAuth.GetUsername(), basicAuth.GetPassword())
	}
	api, err := url.JoinPath(p.c.GetEndpoint(), prometheusAPIV1Metadata)
	if err != nil {
		return nil, err
	}
	getResponse, err := hx.Get(api)
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()
	if getResponse.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(getResponse.Body)
		return nil, merr.ErrorBadRequest("status code: %d => %s", getResponse.StatusCode, string(body))
	}
	var allResp PromMetadataResponse
	if err = json.NewDecoder(getResponse.Body).Decode(&allResp); err != nil {
		return nil, err
	}
	return allResp.Data, nil
}
