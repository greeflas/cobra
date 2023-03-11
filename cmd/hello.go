package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use: "hello",
	//Args: cobra.MinimumNArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing name argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello, %s!\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
