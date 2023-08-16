package main

import "github.com/kaidev0711/go-project/api"

func main() {
	service := api.NewService()

	service.Start()
}
