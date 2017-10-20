package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Follow these directions to create a token:
	// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
	var token = flag.String("token", "", "GitHub api token. Should be restricted to repo:status scope.")
	var targetUrl = flag.String("target_url", "", "Url that this status should redirect to.")
	var statusCtxt = flag.String("context", "status", "Unique string identifier for this status. Something like 'compile', 'test', 'deploy'.")
	var desc = flag.String("description", "", "Description of the test, etc.")
	var state = flag.String("state", "", "Repo status. Can be one of 'error', 'failure', 'pending', 'success'")
	var repo = flag.String("repo", "", "Repository name.")
	var username = flag.String("username", "", "Username")
	var rev = flag.String("rev", "", "Git commit/revision specifier")
	flag.Parse()

	if *state == "" {
		log.Fatal("Please provide a state")
	}
	// if *desc == "" {
	// 	log.Fatal("Please provide a description")
	// }

	// pull from circleci env
	if *token == "" {
		*repo = os.Getenv("GITHUB_API_TOKEN")
	}
	if *repo == "" {
		*repo = os.Getenv("CIRCLE_PROJECT_REPONAME")
	}
	if *rev == "" {
		*rev = os.Getenv("CIRCLE_SHA1")
	}
	if *targetUrl == "" {
		*targetUrl = os.Getenv("CIRCLE_BUILD_URL")
	}
	if *username == "" {
		*username = os.Getenv("CIRCLE_PROJECT_USERNAME")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	status := github.RepoStatus{
		TargetURL:   &*targetUrl,
		State:       &*state,
		Description: &*desc,
		Context:     &*statusCtxt,
	}
	ss, resp, err := client.Repositories.CreateStatus(ctx, *username, *repo, *rev, &status)
	checkErr(err)
	log.Println(resp)
	log.Println(ss)
}

func printJson(x interface{}) {
	b, err := json.Marshal(x)
	checkErr(err)
	os.Stdout.Write(b)
}

func listOrgs(ctx context.Context, client github.Client) {
	orgs, _, err := client.Organizations.List(ctx, "justbuchanan", nil)
	checkErr(err)
	printJson(orgs)
}
