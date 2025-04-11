package vobj

import (
	"fmt"

	"github.com/moon-monitor/moon/pkg/api/houyi/common"
)

func MetricDatasourceUniqueKey(driver common.MetricDatasourceItem_Driver, id uint32) string {
	return fmt.Sprintf("%d:%d", driver, id)
}

func MetricRuleUniqueKey(teamId uint32, strategyId uint32, levelId uint32, datasourceUniqueKey string) string {
	return fmt.Sprintf("%d:%d:%d:%s", teamId, strategyId, levelId, datasourceUniqueKey)
}
