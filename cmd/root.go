package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:     "demo",
	Short:   "This is demo cobra app",
	Long:    "This is demo Cobra application for learning purpose",
	Version: "v1.0.0",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("[%s] ", time.Now().Format("2006-01-02 03:04:05"))
	},
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
