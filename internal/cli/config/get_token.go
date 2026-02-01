package config

import "github.com/spf13/cobra"

var getTokenCmd = &cobra.Command{
	Use:   "get-token",
	Short: "print existing token",
	Run:   func(cmd *cobra.Command, args []string) {},
}
