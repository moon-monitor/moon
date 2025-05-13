package main

import (
	"fmt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"

	"github.com/moon-monitor/moon/cmd/houyi/internal/conf"
	"github.com/moon-monitor/moon/pkg/hello"
	mlog "github.com/moon-monitor/moon/pkg/log"
	"github.com/moon-monitor/moon/pkg/plugin/registry"
	"github.com/moon-monitor/moon/pkg/plugin/server"
	"github.com/moon-monitor/moon/pkg/util/load"
)

// Version is the version of the compiled software.
var Version string
var cfgPath string
var rootCmd = &cobra.Command{
	Use:   "moon",
	Short: "CLI for managing Moon monitor houyi Server",
	Long:  `The Moon Server CLI provides a command-line interface for managing and interacting with the Moon monitor houyi Server service.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the moon houyi service from Moon Monitor!")
		run(cfgPath)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "conf", "c", "./cmd/houyi/config", "Path to the configuration files")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(cfgPath string) {
	var bc conf.Bootstrap
	if err := load.Load(cfgPath, &bc); err != nil {
		panic(err)
	}
	hello.SetEnvWithConfig(Version, bc.GetEnvironment(), bc.GetServer())

	logger, err := mlog.New(bc.IsDev(), bc.GetLog())
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(&bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(c *conf.Bootstrap, srvs server.Servers, logger log.Logger) *kratos.App {
	hello.Hello()
	env := hello.GetEnv()
	opts := []kratos.Option{
		kratos.ID(env.ID()),
		kratos.Name(env.Name()),
		kratos.Version(env.Version()),
		kratos.Metadata(env.Metadata()),
		kratos.Logger(logger),
		kratos.Server(srvs...),
	}
	registerConf := c.GetRegistry()
	if registerConf != nil && registerConf.GetEnable() {
		reg, err := registry.NewRegister(c.GetRegistry())
		if err != nil {
			panic(err)
		}
		opts = append(opts, kratos.Registrar(reg))
	}

	return kratos.New(opts...)
}
