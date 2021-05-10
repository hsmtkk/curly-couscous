package savedata

import (
	"log"

	"github.com/hsmtkk/curly-couscous/database"
	"github.com/hsmtkk/curly-couscous/html2md"
	"github.com/hsmtkk/curly-couscous/htmlparse"
	"github.com/hsmtkk/curly-couscous/http"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "savedata url",
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
	title, err := htmlparse.New().GetTitle(html)
	if err != nil {
		log.Fatal(err)
	}
	md, err := html2md.New().Convert(html)
	if err != nil {
		log.Fatal(err)
	}
	rc := database.Record{
		URL:      url,
		Title:    title,
		MarkDown: md,
	}
	op, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := op.Write(rc); err != nil {
		log.Fatal(err)
	}
}
