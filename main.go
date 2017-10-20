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
	var targetUrl = flag.String("target_url", "https://github.com", "Url that this status should redirect to.")
	var statusCtxt = flag.String("context", "status", "Unique string identifier for this status. Something like 'compile', 'test', 'deploy'.")
	var desc = flag.String("description", "", "Description of the test, etc.")
	var state = flag.String("state", "error", "Repo status. Can be one of 'error', 'failure', 'pending', 'success'")
	flag.Parse()

	if *token == "" {
		log.Fatal("Invalid GitHub status token")
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
	ss, resp, err := client.Repositories.CreateStatus(ctx, "justbuchanan", "ci-test", "8827dd06dfdf43f389b39b90f3d33d0e303bfc2d", &status)
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
