package main

import (
	"log"

	"github.com/kaidev0711/go-project/api"
	"github.com/kaidev0711/go-project/infras/config"
)

func main() {
	var err error

	err = config.StartConfig()
	FatalError(err)
	err = api.NewService().Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
