package hook

import "context"

type Sender interface {
	Send(ctx context.Context, message Message) error
}

type BasicAuth struct {
	Username string
	Password string
}
