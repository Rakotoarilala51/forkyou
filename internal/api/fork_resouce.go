package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Rakotoarilala51/rin"
)

func ForkRepository(repository string) error {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return fmt.Errorf("Repository must be in the format owner/project")
	}
	return GithubAPI().Call("fork", map[string]string{
		"owner": values[0],
		"repo":  values[1],
	}, nil)
}

func ForkSuccess(resp *http.Response) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := ForkResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Println("Forked to repository: %s", respContent.FullName)
	return nil
}
func GetForkRessource() *rin.RestRessources {
	forkRouter := rin.NewRouter()
	forkRouter.RegisterFunc(202, ForkSuccess)
	forkRouter.RegisterFunc(401, func(_ *http.Response) error {
		return fmt.Errorf("You must set an authentication Token")
	})
	fork := rin.NewRessource("/repos/{{.owner}}/{{.repo}}/forks", "POST", forkRouter)
	return fork
}
