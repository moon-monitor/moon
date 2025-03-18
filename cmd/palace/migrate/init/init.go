//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"
	"github.com/spf13/cobra"
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

	mainDB, err := gorm.NewDB(bc.GetData().GetMain())
	if err != nil {
		panic(err)
	}

	if err := mainDB.GetDB().AutoMigrate(system.Models()...); err != nil {
		panic(err)
	}
}
