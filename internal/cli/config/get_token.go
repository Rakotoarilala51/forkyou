package config

import (
	"fmt"

	"github.com/Rakotoarilala51/forkyou/internal/config"
	"github.com/spf13/cobra"
)

var getTokenCmd = &cobra.Command{
	Use:   "get-token",
	Short: "print existing token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.GetToken())
	},
}
