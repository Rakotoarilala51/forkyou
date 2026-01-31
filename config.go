package forkyou

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Managing configuration file of forkyou",
}

func init() {
	ConfigCmd.AddCommand(AddTokenCmd)
}

var AddTokenCmd = &cobra.Command{
	Use:   "add-token",
	Short: "adding your access token in the config file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("token added")
	},
}
