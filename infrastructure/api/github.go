package api

import (
	"context"

	"github.com/google/go-github/v27/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	*github.Client
}

func NewGitHubClient(githubToken string) *GitHubClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &GitHubClient{
		client,
	}
}
