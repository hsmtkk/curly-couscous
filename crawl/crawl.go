package main

import (
	"log"

	"github.com/hsmtkk/curly-couscous/crawl/getweb"
	"github.com/hsmtkk/curly-couscous/crawl/finddata"
	"github.com/hsmtkk/curly-couscous/crawl/savedata"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use: "crawl",
}

func init() {
	command.AddCommand(getweb.Command)
	command.AddCommand(savedata.Command)
	command.AddCommand(finddata.Command)
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
