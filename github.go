package main

import (
	"context"

	"github.com/google/go-github/v27/github"
	"golang.org/x/oauth2"
)

type GitHubAPI struct {
	owner      string
	repository string
	client     *github.Client
	ctx        *context.Context
}

func NewGitHubAPI(owner, repository, githubToken string) *GitHubAPI {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &GitHubAPI{
		owner:      owner,
		repository: repository,
		client:     client,
		ctx:        &ctx,
	}
}

func (g *GitHubAPI) CreateIssueByNumber(title string) (*github.Issue, error) {
	input := &github.IssueRequest{
		Title: github.String(title),
	}

	issue, _, err := g.client.Issues.Create(*g.ctx, g.owner, g.repository, input)
	return issue, err
}
