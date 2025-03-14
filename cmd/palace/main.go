package main

import (
	"fmt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/server"
	"github.com/moon-monitor/moon/pkg/hello"
	mlog "github.com/moon-monitor/moon/pkg/log"
)

// Version is the version of the compiled software.
var Version string
var cfgPath string
var rootCmd = &cobra.Command{
	Use:   "moon",
	Short: "CLI for managing Moon monitor palace Server",
	Long:  `The Iter X Server CLI provides a command-line interface for managing and interacting with the Moon monitor palace Server service.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the moon palace service from Moon Monitor!")
		run(cfgPath)
	},
}

func init() {
	// 增加这段代码
	json.MarshalOptions = protojson.MarshalOptions{
		UseEnumNumbers:  true, // 将枚举值作为数字发出，默认为枚举值的字符串
		UseProtoNames:   true, // 使用 proto 的字段名作为输出字段名
		EmitUnpopulated: true, // 输出未设置字段
	}
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "conf", "c", "./cmd/palace/config", "Path to the configuration files")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(cfgPath string) {
	bc, err := conf.Load(cfgPath)
	if err != nil {
		panic(err)
	}
	logger, err := mlog.New(bc.IsDev(), bc.GetLog())
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(c *conf.Bootstrap, srv *server.Server, logger log.Logger) *kratos.App {
	envOpts := []hello.Option{
		hello.WithVersion(Version),
		hello.WithName("Palace"),
	}
	hello.SetEnvWithOption(envOpts...)
	hello.Hello()
	env := hello.GetEnv()
	opts := []kratos.Option{
		kratos.ID(env.ID()),
		kratos.Name(env.Name()),
		kratos.Version(env.Version()),
		kratos.Metadata(env.Metadata()),
		kratos.Logger(logger),
		kratos.Server(srv.GetServers()...),
	}
	//registerConf := c.GetDiscovery()
	//if !types.IsNil(registerConf) {
	//	register, err := conn.NewRegister(c.GetDiscovery(), conn.WithDiscoveryConfigEtcd(c.GetDiscovery().GetEtcd()))
	//	if !types.IsNil(err) {
	//		log.Warnw("register error", err)
	//		panic(err)
	//	}
	//	opts = append(opts, kratos.Registrar(register))
	//}

	return kratos.New(opts...)
}
