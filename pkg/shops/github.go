package shops

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func commits(ctx context.Context) ([]string, error) {
	gc, err := newGithubClient(ctx)
	if err != nil {
		return nil, err
	}
	opts := &github.CommitsListOptions{
		SHA: "master",
		ListOptions: github.ListOptions{
			PerPage: 10,
		},
	}
	commits, _, err := gc.Repositories.ListCommits(ctx, "kastenhq", "k10", opts)
	if err != nil {
		return nil, err
	}
	cs := make([]string, 0, len(commits))
	for _, c := range commits {
		//fmt.Printf("Sha: %#v\n", (*c.SHA)[:8])
		cs = append(cs, *c.SHA)
	}
	return cs, nil
}

const githubTokenEnvVar = `GITHUB_TOKEN`

func newGithubClient(ctx context.Context) (*github.Client, error) {
	token, ok := os.LookupEnv(githubTokenEnvVar)
	if !ok {
		return nil, fmt.Errorf("please set $%s", githubTokenEnvVar)
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), nil
}
