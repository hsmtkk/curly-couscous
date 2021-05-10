package cmd

import (
	"github.com/hsmtkk/curly-couscous/cmd/getweb"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "crawl",
}

func init() {
	Command.AddCommand(getweb.Command)
}
