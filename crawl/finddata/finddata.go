package finddata

import(
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/hsmtkk/curly-couscous/database"
)

var Command = &cobra.Command {
	Use: "finddata title",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		run(args[0])
	},
}

func run(title string){
	op, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	rcd, err := op.Find(title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%#v", rcd)
}