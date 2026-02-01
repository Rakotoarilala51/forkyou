package config

import (
	"log"

	"github.com/Rakotoarilala51/forkyou/internal/config"
	"github.com/spf13/cobra"
)

var addTokenCmd = &cobra.Command{
	Use:   "add-token",
	Short: "adding your access token in the config file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("you must supply the token bro")
		}
		config.AddToken(args[0])
	},
}
