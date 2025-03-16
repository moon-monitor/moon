package vobj

// OAuthAPP oauth app
//
//go:generate stringer -type=OAuthAPP -linecomment -output=oauth_app.string.go
type OAuthAPP int8

const (
	OAuthAPPGithub OAuthAPP = iota // github
	OAuthAPPGitee                  // gitee
)
