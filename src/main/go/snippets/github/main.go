package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pwera/github/github"
)

func main() {
	token, _ := os.LookupEnv("token")
	background := context.Background()
	user := "pwera"
	service := github.RepositoriesService{Client: github.NewClient(background, token)}
	list, _, err := service.List(background, user)
	if err != nil {
		log.Print(err)
	}
	for i, r := range list {
		fmt.Printf("%d: %v\n", i, r)
	}
}
