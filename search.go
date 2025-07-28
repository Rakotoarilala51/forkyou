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

type SearchResponse struct{
	Results []*SearchResult `json:"items"`
}
type SearchResult struct{
	FullName string `json:"full_name"`
}

var SearchCmd = &cobra.Command{
	Use: "search",
	Short: "search for a github  repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if err := SearchByKeyword(args); err != nil{
			log.Fatalln("Search Failed:", err)
		}
	},
}

func SearchByKeyword(keywords []string) error{
	return GithubAPI().Call("search", map[string]string{
		"query":"?q="+strings.Join(keywords, "+"),
	}, nil)
}

func SearchSuccess(resp *http.Response) error{
	defer resp.Body.Close()
	content , err:= ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	respContent := SearchResponse{}
	json.Unmarshal(content, &respContent)
	for _, item := range respContent.Results{
		fmt.Println(item.FullName)
	}
	return nil
}

func SearchDefaultRouter(resp *http.Response) error{
	return fmt.Errorf("Status code %d", resp.StatusCode)
}

func GetSearchRessource() *rin.RestRessources{
	searchRouter := rin.NewRouter()
	searchRouter.DefaultRouter = SearchDefaultRouter
	searchRouter.RegisterFunc(200, SearchSuccess)
	search := rin.NewRessource("/search/repositories{{.query}}", "GET", searchRouter)
	return search
}