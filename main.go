package main

import (
	"os"
)

func main() {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	botID := os.Getenv("BOT_ID")
	channelID := os.Getenv("CHANNEL_ID")
	githubToken := os.Getenv("GITHUB_TOKEN")
	githubOwner := os.Getenv("GITHUB_OWNER")
	githubRepository := os.Getenv("GITHUB_REPOSITORY")

	githubAPI := NewGitHubAPI(githubOwner, githubRepository, githubToken)
	bot := NewBot(botID, channelID)
	messenger := NewMessenger(apiToken, bot, githubAPI)
	messenger.Listen()
}
