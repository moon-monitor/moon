package team

func Models() []any {
	return []any{
		&Dashboard{},
		&DashboardChart{},
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
		&TimeEngine{},
		&TimeEngineRule{},
	}
}
