package main

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	client := github.NewClient(nil)

	orgs, _, err := client.Organizations.List(ctx, "justbuchanan", nil)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(orgs)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)
}
