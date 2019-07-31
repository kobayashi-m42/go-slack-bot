package repository

import (
	"github.com/kobayashi-m42/go-slack-bot/domain"
)

type GitHubRepository interface {
	CreateIssue(p *domain.CreateIssue) (*domain.Issue, error)
}
