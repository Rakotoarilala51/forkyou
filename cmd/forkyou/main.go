package main

import (
	"github.com/Rakotoarilala51/forkyou/internal/cli"
	"github.com/Rakotoarilala51/forkyou/internal/config"
)

func main() {
	config.InitConfig()
	cli.Execute()
}
