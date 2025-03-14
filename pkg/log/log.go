package log

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/pkg/config"
)

func New(isDev bool, cfg *config.Log) (log.Logger, error) {
	switch cfg.GetDriver() {
	case config.Log_SUGARED:
		return NewSugaredLogger(isDev, cfg.GetLevel(), cfg.GetSugared())
	default:
		return log.NewStdLogger(os.Stdout), nil
	}
}
