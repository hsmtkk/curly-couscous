package getweb

import (
	"fmt"
	"log"

	"github.com/hsmtkk/curly-couscous/html2md"
	"github.com/hsmtkk/curly-couscous/http"
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
	html, err := http.New().Get(url)
	if err != nil {
		log.Fatal(err)
	}
	md, err := html2md.New().Convert(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(md)
}
