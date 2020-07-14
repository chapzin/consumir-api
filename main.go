package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/chapzin/consumir-api/client"
)

func main() {
	c := client.NewClient(nil)

	owner := flag.String("owner", "golang", "name of the repository owner")
	repoName := flag.String("repo", "go", "name of the repository")

	flag.Parse()

	repo, _, err := c.Repository.Get(*owner, *repoName)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Name: %s \nDescription: %s\nURL: %s\nStarts: %d\nWatchers: %d\n", repo.FullName, repo.Description, repo.HTMLURL, repo.Starts, repo.Watchers)
}
