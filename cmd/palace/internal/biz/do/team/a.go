package team

func Models() []any {
	return []any{
		&EmailConfig{},
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
