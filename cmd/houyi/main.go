package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is the version of the compiled software.
var Version string
var cfgPath string
var rootCmd = &cobra.Command{
	Use:   "moon",
	Short: "CLI for managing Moon monitor houyi Server",
	Long:  `The Iter X Server CLI provides a command-line interface for managing and interacting with the Moon monitor houyi Server service.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the moon houyi service from Moon Monitor!")
		run(cfgPath)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "conf", "c", "./cmd/houyi/config", "Path to the configuration files")
}

func main() {
	fmt.Println("Welcome to the moon 后裔 service from Moon Monitor!")
}

func run(cfgPath string) {}
