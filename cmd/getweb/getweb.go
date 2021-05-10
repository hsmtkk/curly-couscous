package getweb

import (
	"fmt"
	"log"

	"github.com/hsmtkk/curly-couscous/getweb"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:  "getweb url",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		run(args[0])
	},
}

func run(url string) {
	getter := getweb.New()
	md, err := getter.GetWeb(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(md)
}
