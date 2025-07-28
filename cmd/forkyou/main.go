package main

import (
	"os"

	"github.com/Rakotoarilala51/forkyou"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("forkyou")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}
