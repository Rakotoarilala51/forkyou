package forkyou

import "github.com/spf13/cobra"

var ForkCmd = &cobra.Command{
	Use: "fork",
	Short: "Forking a github repository",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}