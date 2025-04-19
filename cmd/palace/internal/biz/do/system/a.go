package system

func Models() []any {
	return []any{
		&OperateLog{},
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
