package main

import (
	"os"
)

func main() {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	botID := os.Getenv("BOT_ID")
	channelID := os.Getenv("CHANNEL_ID")

	bot := NewBot(botID, channelID)
	messenger := NewMessenger(apiToken, bot)
	messenger.Listen()
}
