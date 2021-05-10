package main

import (
	"log"

	"github.com/hsmtkk/curly-couscous/cmd"
)

func main() {
	c := cmd.Command
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}
