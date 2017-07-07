package internal

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func NewGithubClient(token *string) *github.Client {
	ctx := context.Background()
	var tc *http.Client

	if token != nil {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: *token},
		)
		tc = oauth2.NewClient(ctx, ts)

	}

	return github.NewClient(tc)
}
