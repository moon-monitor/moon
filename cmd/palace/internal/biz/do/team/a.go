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
