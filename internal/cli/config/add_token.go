package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddTokenCmd = &cobra.Command{
	Use:   "add-token",
	Short: "adding your access token in the config file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("token added")
	},
}
