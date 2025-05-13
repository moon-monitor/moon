package hello

import (
	"os"
	"sync"

	"github.com/moon-monitor/moon/pkg/config"
)

var (
	once  sync.Once
	id, _ = os.Hostname()
	env   = &Env{
		name:     "Moon",
		version:  "v3.0.0",
		metadata: make(map[string]string),
		id:       id,
		env:      config.Environment_PROD.String(),
	}
)

type Env struct {
	name     string
	version  string
	metadata map[string]string
	id       string
	env      string
}

type Option func(*Env)

func (e *Env) Name() string {
	return env.name
}

func (e *Env) Version() string {
	return env.version
}

func (e *Env) Metadata() map[string]string {
	return env.metadata
}

func (e *Env) ID() string {
	return env.id
}

func (e *Env) Env() string {
	return env.env
}

func SetEnvWithOption(opts ...Option) {
	once.Do(func() {
		for _, opt := range opts {
			opt(env)
		}
	})
}

func SetEnvWithConfig(version string, env config.Environment, serverConf *config.Server) {
	opts := []Option{
		WithVersion(version),
		WithEnv(env),
		WithName(serverConf.GetName()),
		WithMetadata(serverConf.GetMetadata()),
	}
	SetEnvWithOption(opts...)
}

func GetEnv() *Env {
	return env
}

func WithName(name string) Option {
	return func(e *Env) {
		e.name = name
	}
}

func WithVersion(version string) Option {
	return func(e *Env) {
		if version != "" {
			e.version = version
		}
	}
}

func WithMetadata(metadata map[string]string) Option {
	return func(e *Env) {
		e.metadata = metadata
	}
}

func WithID(id string) Option {
	return func(e *Env) {
		e.id = id
	}
}

func WithEnv(env config.Environment) Option {
	return func(e *Env) {
		e.env = env.String()
	}
}
