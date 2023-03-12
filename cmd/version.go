package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var short bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			fmt.Println("v1.0.0")
		} else {
			fmt.Println("Cobra Demo App v1.0.0")
		}
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&short, "short", "s", false, "print short version")
	RootCmd.AddCommand(versionCmd)
}
