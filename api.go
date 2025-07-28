package forkyou

import (
	"github.com/Rakotoarilala51/rin"
	"github.com/spf13/viper"
)

var api *rin.API

func GithubAPI() *rin.API {
	if api ==nil{
		api = rin.NewApi("https://api.github.com")
		token := viper.GetString("token")
		api.SetAuth(rin.NewAuthToken(token))
		api.AddRessource("fork", GetForkRessource())
		api.AddRessource("search", GetSearchRessource())
	}
	return api
}
