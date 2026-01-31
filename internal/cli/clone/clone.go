package clone

import (
	"fmt"
	"log"

	"github.com/Rakotoarilala51/forkyou/internal/core/repo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "cloning git repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply repository name")
		}
		if err := CloneRepository(args[0], ref, create); err != nil {
			log.Fatalln("Error when cloning repository:", err)
		}
	},
}

var ref string
var create bool

func CloneRepository(repository string, ref string, shouldCreate bool) error {
	repoGh, err := repo.NewGHRepo(repository)
	if err != nil {
		return err
	}

	cloneService := repo.NewCloneService(repoGh)
	if err := cloneService.Execute(viper.GetString("location"), ref, shouldCreate); err != nil {
		return err
	}

	fmt.Println("Repository cloned successfully")
	return nil
}
func init() {
	CloneCmd.PersistentFlags().StringVar(&ref, "ref", "master", "specific reference to checkout")
	CloneCmd.PersistentFlags().BoolVar(&create, "create", false, "create the reference if it doesn't exist")
}
