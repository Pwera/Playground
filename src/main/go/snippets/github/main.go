package main

import (
	"context"
	"github.com/pwera/Playground/src/main/go/snippets/github/github"
)

func main() {
	background := context.Background()
	user := "pwera"
	token := "//"
	service := github.RepositoriesService{client: NewClient(background, token)}
	service.List(background, user)
}
