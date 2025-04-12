package system

func Models() []any {
	return []any{
		&Menu{},
		&Resource{},
		&Role{},
		&Team{},
		&TeamAudit{},
		&User{},
		&UserOAuth{},
	}
}
