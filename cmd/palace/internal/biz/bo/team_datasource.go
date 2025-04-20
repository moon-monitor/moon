package bo

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

type SaveTeamMetricDatasource struct {
	ID             uint32
	Name           string
	Status         vobj.GlobalStatus
	Remark         string
	Driver         vobj.DatasourceDriverMetric
	Endpoint       string
	ScrapeInterval time.Duration
	Headers        kv.StringMap
	QueryMethod    vobj.HTTPMethod
	CA             string
	TLS            *do.TLS
	BasicAuth      *do.BasicAuth
	Extra          kv.StringMap
}

type ListTeamMetricDatasource struct {
	*PaginationRequest
	Status  vobj.GlobalStatus
	Keyword string
}

func (r *ListTeamMetricDatasource) ToListTeamMetricDatasourceReply(datasourceItems []*team.DatasourceMetric) *ListTeamMetricDatasourceReply {
	return &ListTeamMetricDatasourceReply{
		PaginationReply: r.ToReply(),
		Items:           slices.Map(datasourceItems, func(datasource *team.DatasourceMetric) do.DatasourceMetric { return datasource }),
	}
}

type ListTeamMetricDatasourceReply = ListReply[do.DatasourceMetric]

type UpdateTeamMetricDatasourceStatusRequest struct {
	DatasourceID uint32
	Status       vobj.GlobalStatus
}
