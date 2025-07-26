package gogit

import "github.com/spf13/cobra"

var SearchCmd = &cobra.Command{
	Use: "search",
	Short: "search for a github  repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}