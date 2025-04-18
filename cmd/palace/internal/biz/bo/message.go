package bo

import (
	"context"
)

type SendEmailFun func(ctx context.Context, params *SendEmailParams) error
