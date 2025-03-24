package team

func Models() []any {
	return []any{
		&Dict{},
		&StrategyGroup{},
		&Strategy{},
		&NoticeGroup{},
		&NoticeHook{},
		&TimeEngine{},
		&TimeEngineRule{},
	}
}
