package prometheus_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/pkg/plugin/datasource"
	"github.com/moon-monitor/moon/pkg/plugin/datasource/prometheus"
)

var _ prometheus.Config = (*config)(nil)

type config struct {
	Endpoint  string
	BasicAuth datasource.BasicAuth
}

func (c *config) GetEndpoint() string {
	return c.Endpoint
}

func (c *config) GetBasicAuth() datasource.BasicAuth {
	return c.BasicAuth
}

func Test_Query(t *testing.T) {
	c := &config{
		Endpoint: "https://prometheus.aide-cloud.cn/",
	}
	prom := prometheus.New(c, log.GetLogger())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	expr := `rate(process_cpu_seconds_total[1m])`
	resp, err := prom.Query(ctx, &datasource.MetricQueryRequest{
		Expr:      expr,
		Time:      time.Now().Add(-time.Second * 30).Unix(),
		StartTime: 0,
		EndTime:   0,
		Step:      0,
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	if !resp.IsSuccessResponse() {
		t.Fatal(resp.Error())
	}

	for _, v := range resp.Data.Result {
		t.Log(v.GetMetricQueryValue())
	}
}

func Test_QueryRange(t *testing.T) {
	c := &config{
		Endpoint: "https://prometheus.aide-cloud.cn/",
	}
	prom := prometheus.New(c, log.GetLogger())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	expr := `rate(process_cpu_seconds_total[1m])`
	now := time.Now()
	resp, err := prom.Query(ctx, &datasource.MetricQueryRequest{
		Expr:      expr,
		Time:      0,
		StartTime: now.Add(-time.Second * 30).Unix(),
		EndTime:   now.Unix(),
		Step:      14,
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	if !resp.IsSuccessResponse() {
		t.Fatal(resp.Error())
	}

	for _, v := range resp.Data.Result {
		t.Log(v.GetMetricQueryValues())
	}
	t.Log(resp.String())
}

func Test_Metadata(t *testing.T) {
	c := &config{
		Endpoint: "https://prometheus.aide-cloud.cn/",
	}
	prom := prometheus.New(c, log.GetLogger())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	metadata, err := prom.Metadata(ctx)
	if err != nil {
		t.Fatal(err)
		return
	}
	for item := range metadata {
		t.Log(item)
	}
}
