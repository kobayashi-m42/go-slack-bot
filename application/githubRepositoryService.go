package application

import (
	"github.com/kobayashi-m42/go-slack-bot/domain"
	"github.com/kobayashi-m42/go-slack-bot/domain/repository"
	"github.com/kobayashi-m42/go-slack-bot/domain/service"
)

type githubService struct {
	githubRepository repository.GitHubRepository
}

func NewGitHubService(r repository.GitHubRepository) service.GitHubRepositoryService {
	return &githubService{
		githubRepository: r,
	}
}

func (g *githubService) CreateIssue(title string) *domain.BotResponse {

	createIssue := &domain.CreateIssue{
		Title: title,
	}

	issue, err := g.githubRepository.CreateIssue(createIssue)
	if err != nil {
		return &domain.BotResponse{
			Msg: "",
			Err: &domain.CreateIssueError{Msg: err.Error()},
		}
	}

	return &domain.BotResponse{
		Msg: domain.CreateIssueSuccessMessage + issue.URL,
		Err: nil,
	}
}
