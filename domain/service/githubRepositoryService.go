package service

import (
	"github.com/kobayashi-m42/go-slack-bot/domain"
)

type GitHubRepositoryService interface {
	CreateIssue(title string) *domain.BotResponse
}
