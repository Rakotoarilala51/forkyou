package search

import (
	"log"

	"github.com/Rakotoarilala51/forkyou/internal/api"
	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for a github  repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if err := api.SearchByKeyword(args); err != nil {
			log.Fatalln("Search Failed:", err)
		}
	},
}
