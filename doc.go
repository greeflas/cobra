package main

import (
	"github.com/greeflas/cobra/cmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {
	if err := doc.GenMarkdownTree(cmd.RootCmd, "./doc"); err != nil {
		log.Fatal(err)
	}
}
