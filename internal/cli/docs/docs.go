package docs

import (
	"fmt"
	"log"

	"github.com/Rakotoarilala51/forkyou/internal/api"
	"github.com/spf13/cobra"
)

var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "read documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must Supply repository argument")
		}
		if err := api.GetRepositoryReadme(args[0]); err != nil {
			fmt.Println("Error when reading docs: ", err)
		}

	},
}
