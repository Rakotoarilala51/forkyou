package cli

import (
	"github.com/Rakotoarilala51/forkyou/internal/cli/clone"
	"github.com/Rakotoarilala51/forkyou/internal/cli/config"
	"github.com/Rakotoarilala51/forkyou/internal/cli/docs"
	"github.com/Rakotoarilala51/forkyou/internal/cli/fork"
	"github.com/Rakotoarilala51/forkyou/internal/cli/search"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "forkyou",
	Short: "CLI tool to fork repo on github",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	RootCmd.AddCommand(
		fork.ForkCmd,
		clone.CloneCmd,
		search.SearchCmd,
		docs.DocsCmd,
		config.ConfigCmd,
	)
}
