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
		&SendMessageLog{},
		&Strategy{},
		&StrategyGroup{},
		&StrategyMetric{},
		&StrategySubscriber{},
		&StrategyMetricRule{},
		&StrategyMetricRuleLabelNotice{},
		&TimeEngine{},
		&TimeEngineRule{},
	}
}
