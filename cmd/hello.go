package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use: "hello",
	SuggestFor: []string{
		"hi",
	},
	//Args: cobra.MinimumNArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing name argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		msg := fmt.Sprintf("Hello, %s", args[0])

		if Verbose {
			msg += "!"
		}

		fmt.Println(msg)
	},
}

func init() {
	RootCmd.AddCommand(helloCmd)
}
