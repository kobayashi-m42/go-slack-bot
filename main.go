package main

import (
	"os"

	"github.com/kobayashi-m42/go-slack-bot/application"
	"github.com/kobayashi-m42/go-slack-bot/handler"
	"github.com/kobayashi-m42/go-slack-bot/infrastructure"
	"github.com/kobayashi-m42/go-slack-bot/infrastructure/api"
	"github.com/kobayashi-m42/go-slack-bot/infrastructure/log"
)

func main() {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	botID := os.Getenv("BOT_ID")
	channelID := os.Getenv("CHANNEL_ID")
	githubToken := os.Getenv("GITHUB_TOKEN")
	githubOwner := os.Getenv("GITHUB_OWNER")
	githubRepository := os.Getenv("GITHUB_REPOSITORY")

	// Initialize GitHub client
	githubClient := api.NewGitHubClient(githubToken)

	// Initialize repository
	repository := infrastructure.NewGitHubRepository(githubOwner, githubRepository, githubClient)

	// Initialize application service
	service := application.NewGitHubService(repository)

	// Initialize Logger
	logger := log.NewLogger()

	// Run Bot
	botServer := handler.NewBot(
		apiToken,
		botID,
		channelID,
		logger,
		&handler.Services{
			GitHubRepositoryService: service,
		})
	botServer.Listen()
}
