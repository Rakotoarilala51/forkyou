package main

import (
	"github.com/Rakotoarilala51/forkyou"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func main(){
	rootCmd.Execute()
}
func init(){
	rootCmd  = &cobra.Command{
		Use: "forkyou",
		Short: "Project Forking Tool for github",
	}
	rootCmd.AddCommand(forkyou.SearchCmd)
	rootCmd.AddCommand(forkyou.DocsCmd)
	rootCmd.AddCommand(forkyou.CloneCmd)
	rootCmd.AddCommand(forkyou.ForkCmd)
}
