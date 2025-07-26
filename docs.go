package gogit

import "github.com/spf13/cobra"

var DocsCmd = &cobra.Command{
	Use: "docs",
	Short: "read documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {},
}