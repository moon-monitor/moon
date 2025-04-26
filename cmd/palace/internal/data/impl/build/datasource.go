package build

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/pkg/util/crypto"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToDatasourceMetric(ctx context.Context, datasource do.DatasourceMetric) *team.DatasourceMetric {
	if validate.IsNil(datasource) {
		return nil
	}
	if datasource, ok := datasource.(*team.DatasourceMetric); ok {
		datasource.WithContext(ctx)
		return datasource
	}
	return &team.DatasourceMetric{
		TeamModel:      ToTeamModel(ctx, datasource),
		Name:           datasource.GetName(),
		Status:         datasource.GetStatus(),
		Remark:         datasource.GetRemark(),
		Driver:         datasource.GetDriver(),
		Endpoint:       datasource.GetEndpoint(),
		ScrapeInterval: datasource.GetScrapeInterval(),
		Headers:        datasource.GetHeaders(),
		QueryMethod:    datasource.GetQueryMethod(),
		CA:             datasource.GetCA(),
		TLS:            crypto.NewObject(datasource.GetTLS()),
		BasicAuth:      crypto.NewObject(datasource.GetBasicAuth()),
		Extra:          datasource.GetExtra(),
		Metrics:        []*team.StrategyMetric{},
	}
}
func ToDatasourceMetrics(ctx context.Context, datasourceList []do.DatasourceMetric) []*team.DatasourceMetric {
	return slices.MapFilter(datasourceList, func(v do.DatasourceMetric) (*team.DatasourceMetric, bool) {
		if validate.IsNil(v) {
			return nil, false
		}
		return ToDatasourceMetric(ctx, v), true
	})
}
