package config

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Managing configuration file of forkyou",
}

func init() {
	ConfigCmd.AddCommand(AddTokenCmd)
}
