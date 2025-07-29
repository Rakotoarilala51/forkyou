package forkyou

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Rakotoarilala51/rin"
	"github.com/spf13/cobra"
)

type ForkResponse struct{
	CloneURL string `json:"clone_url"`
	FullName string `json:"full_name"`
}

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
	values := strings.Split(repository, "/")
	if len(values) !=2{
		return fmt.Errorf("Repository must be in the format owner/project")
	}
	return GithubAPI().Call("fork", map[string]string{
		"owner":values[0],
		"repo":values[1],
	}, nil)
}

func ForkSuccess(resp *http.Response) error{
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		return err
	}
	respContent := ForkResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Println("Forked to repository: %s", respContent.FullName)
	return nil
}
func GetForkRessource() *rin.RestRessources{
	forkRouter := rin.NewRouter()
	forkRouter.RegisterFunc(202, ForkSuccess)
	forkRouter.RegisterFunc(401, func(_ *http.Response) error {
		return fmt.Errorf("You must set an authentication Token")
	})
	fork := rin.NewRessource("/repos/{{.owner}}/{{.repo}}/forks", "POST", forkRouter)
	return fork
}