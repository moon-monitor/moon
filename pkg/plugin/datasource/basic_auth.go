package datasource

type BasicAuth interface {
	GetUsername() string
	GetPassword() string
}
