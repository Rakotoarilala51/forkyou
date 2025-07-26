package main

import (
	"github.com/Rakotoarilala51/gogit"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func main(){
	rootCmd.Execute()
}
func init(){
	rootCmd  = &cobra.Command{
		Use: "gogit",
		Short: "Project Forking Tool for github",
	}
	rootCmd.AddCommand(gogit.SearchCmd)
}
