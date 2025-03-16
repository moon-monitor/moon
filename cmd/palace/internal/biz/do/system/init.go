package system

func Models() []any {
	return []any{
		&User{},
		&Team{},
		&TeamRole{},
		&TeamMember{},
		&Role{},
		&Resource{},
		&OAuthUser{},
	}
}
