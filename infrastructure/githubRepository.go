package infrastructure

import (
	"context"

	"github.com/kobayashi-m42/go-slack-bot/domain"
	"github.com/kobayashi-m42/go-slack-bot/domain/repository"
	"github.com/kobayashi-m42/go-slack-bot/infrastructure/api"

	"github.com/google/go-github/v27/github"
)

type githubRepository struct {
	owner      string
	repository string
	client     *api.GitHubClient
}

func NewGitHubRepository(owner, repository string, client *api.GitHubClient) repository.GitHubRepository {
	return &githubRepository{
		owner:      owner,
		repository: repository,
		client:     client,
	}
}

func (g *githubRepository) CreateIssue(i *domain.CreateIssue) (*domain.Issue, error) {
	issueRequest := &github.IssueRequest{
		Title: github.String(i.Title),
	}

	ctx := context.Background()
	result, _, err := g.client.Issues.Create(ctx, g.owner, g.repository, issueRequest)

	if err != nil {
		return &domain.Issue{}, err
	}

	issue := &domain.Issue{
		Title: result.GetTitle(),
		URL:   result.GetHTMLURL(),
	}

	return issue, err
}
