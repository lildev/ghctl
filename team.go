package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/google/go-github/v28/github"
)

func getTeams(org string, format string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	opt := &github.ListOptions{PerPage: 100}
	var objsAll []*github.Team
	for {
		objs, resp, err := client.Teams.ListTeams(ctx, org, opt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		objsAll = append(objsAll, objs...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	sort.Slice(objsAll, func(i, j int) bool {
		return *objsAll[i].Name < *objsAll[j].Name
	})

	if format == "normal" {
		for _, repo := range objsAll {
			fmt.Println(*repo.Name)
		}
	} else if format == "json" {
		bytes, _ := json.Marshal(objsAll)
		fmt.Println(string(bytes))
	}
}
