package system

func Models() []any {
	return []any{
		&User{},
		&Team{},
		&TeamRole{},
		&TeamMember{},
		&SysRole{},
		&Resource{},
		&OAuthUser{},
		&TeamAudit{},
		&Menu{},
	}
}
