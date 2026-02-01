package api

import (
	"github.com/Rakotoarilala51/forkyou/internal/config"
	"github.com/Rakotoarilala51/rin"
)

var api *rin.API

func GithubAPI() *rin.API {
	if api == nil {
		api = rin.NewApi("https://api.github.com")
		token := config.GetToken()
		api.SetAuth(rin.NewAuthToken(token))
		api.AddRessource("fork", GetForkRessource())
		api.AddRessource("search", GetSearchRessource())
		api.AddRessource("docs", GetReadmeRessource())
	}
	return api
}
