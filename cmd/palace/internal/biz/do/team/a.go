package team

func Models() []any {
	return []any{
		&SmsConfig{},
		&EmailConfig{},
		&Dashboard{},
		&DashboardChart{},
		&DatasourceMetric{},
		&Dict{},
		&OperateLog{},
		&NoticeGroup{},
		&NoticeHook{},
		&NoticeMember{},
		&Strategy{},
		&StrategyGroup{},
		&StrategyMetric{},
		&StrategyMetricRule{},
		&StrategyMetricRuleLabelNotice{},
		&TimeEngine{},
		&TimeEngineRule{},
	}
}
