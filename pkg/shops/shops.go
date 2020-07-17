package shops

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func Shops(ctx context.Context) error {
	gc, err := newGithubClient(ctx)
	if err != nil {
		return err
	}
	return nil
}

const tokenEnvVar = `GITHUB_TOKEN`

func newGithubClient(ctx context.Context) (*github.Client, error) {
	token, ok := os.LookupEnv(tokenEnvVar)
	if !ok {
		return nil, fmt.Errorf("please set $%s", tokenEnvVar)
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), nil
}
