package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "This is demo cobra app",
	Long:  "This is demo Cobra application for learning purpose",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
