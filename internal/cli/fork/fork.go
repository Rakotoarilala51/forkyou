package fork

import (
	"log"

	"github.com/Rakotoarilala51/forkyou/internal/api"
	"github.com/spf13/cobra"
)

var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Forking repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("You must supply repository name")
		}
		if err := api.ForkRepository(args[0]); err != nil {
			log.Fatalln("Unable to fork Repository")
		}
	},
}
