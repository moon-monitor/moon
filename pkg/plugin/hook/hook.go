package hook

import "context"

type Hook interface {
	Send(ctx context.Context, message Message) error
}

type BasicAuth struct {
	Username string
	Password string
}
