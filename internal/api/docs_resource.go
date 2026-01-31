package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Rakotoarilala51/rin"
)

func GetRepositoryReadme(repository string) error {
	values := strings.Split(repository, "/")
	return GithubAPI().Call("docs", map[string]string{
		"owner":   values[0],
		"project": values[1],
	}, nil)
}
func ReadmeSuccess(resp *http.Response) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := ReadmeResponse{}
	json.Unmarshal(content, &respContent)
	buff, err := base64.StdEncoding.DecodeString(respContent.Content)
	if err != nil {
		return err
	}
	fmt.Println(string(buff))
	return nil
}

func ReadmeDefaultRouter(resp *http.Response) error {
	return fmt.Errorf("status code: %d", resp.StatusCode)
}

func GetReadmeRessource() *rin.RestRessources {
	router := rin.NewRouter()
	router.RegisterFunc(200, ReadmeSuccess)
	router.DefaultRouter = ReadmeDefaultRouter
	ressource := rin.NewRessource("/repos/{{.owner}}/{{.project}}/readme", "GET", router)
	return ressource
}
