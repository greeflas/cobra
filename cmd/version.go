package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cobra Demo App v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
