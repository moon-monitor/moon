package team

func Models() []any {
	return []any{
		&Dashboard{},
		&DashboardChart{},
		&DatasourceMetric{},
		&Dict{},
		&Member{},
		&Menu{},
		&NoticeGroup{},
		&NoticeHook{},
		&NoticeMember{},
		&Resource{},
		&Role{},
		&Strategy{},
		&StrategyGroup{},
		&StrategyMetric{},
		&StrategyMetricRule{},
		&StrategyMetricRuleLabelNotice{},
		&TimeEngine{},
		&TimeEngineRule{},
	}
}
