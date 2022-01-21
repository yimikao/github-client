package main

import (
	"githubClient/client"
)

func main() {
	c := client.NewClient(1)

	client.CreateGist(c, &client.CreateGistParams{})
}
