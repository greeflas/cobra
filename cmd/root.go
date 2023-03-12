package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:     "demo",
	Short:   "This is demo cobra app",
	Long:    "This is demo Cobra application for learning purpose",
	Version: "v1.0.0",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
