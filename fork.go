package forkyou

import (
	"log"

	"github.com/spf13/cobra"
)

var ForkCmd = &cobra.Command{
	Use: "fork",
	Short: "Forking a github repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)<=0{
			log.Fatalln("You must supply repository")
		}
		if err := ForkRepository(args[0]); err != nil{
			log.Fatalln("Unable to fork Repository", err)
		}
	},
}

func ForkRepository(repository string) error{
	return nil
}