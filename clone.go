package forkyou

import (
	"log"

	"github.com/spf13/cobra"
)

var CloneCmd = &cobra.Command{
	Use: "clone",
	Short: "cloning git repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)<=0 {
			log.Fatalln("You must supply repository name")
		}
		if err:=CloneRepository(args[0]); err!=nil{
			log.Fatalln("Error when cloning repository:", err)
		}
	},
}

var ref string
var create bool

func CloneRepository(repository string) error{
	return nil
}

func init(){
	CloneCmd.PersistentFlags().StringVar(&ref, "ref", "","specific reference to checkout")
	CloneCmd.PersistentFlags().BoolVar(&create, "create", false,"create the reference if it doesn't exist")
}