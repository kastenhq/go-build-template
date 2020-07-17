package myapp

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v23/github"
	"golang.org/x/oauth2"
)

func Foo(ctx context.Context) error {
	return nil
}

const tokenEnvVar = `GITHUB_TOKEN`

func newClient(ctx context.Context) (*github.Client, error) {
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
