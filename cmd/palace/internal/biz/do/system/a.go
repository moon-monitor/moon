package system

func Models() []any {
	return []any{
		&Menu{},
		&Resource{},
		&Role{},
		&Team{},
		&TeamAudit{},
		&TeamInviteLink{},
		&TeamInviteUser{},
		&User{},
		&UserConfigTable{},
		&UserConfigTheme{},
		&UserOAuth{},
	}
}
