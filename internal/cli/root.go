package cli

import (
	"github.com/Rakotoarilala51/forkyou/internal/cli/clone"
	"github.com/Rakotoarilala51/forkyou/internal/cli/fork"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forkyou",
	Short: "CLI tool to fork repo on github",
}

func init() {
	rootCmd.AddCommand(
		fork.ForkCmd,
		clone.CloneCmd,
	)
}
