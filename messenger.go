package main

import (
	"fmt"
	"log"

	"github.com/nlopes/slack"
)

type Messenger interface {
	Listen()
	sndMessage(text string, channelID string)
}

type slackMessenger struct {
	rtm *slack.RTM
	Bot
	GitHubAPI
}

func NewMessenger(token string, bot *Bot, githubAPI *GitHubAPI) Messenger {
	api := slack.New(token)
	rtm := api.NewRTM()

	return &slackMessenger{
		rtm:       rtm,
		Bot:       *bot,
		GitHubAPI: *githubAPI,
	}
}

func (m *slackMessenger) Listen() {
	go m.rtm.ManageConnection()

	for msg := range m.rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := m.ValidateMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle responseMessage: %s", err)
				break
			}

			title := m.GetTitleFromText(ev.Msg.Text)
			issue, err := m.CreateIssueByNumber(title)
			if err != nil {
				log.Printf("[ERROR] Failed to create Issue: %s", err)
				break
			}

			text := issue.GetHTMLURL()
			m.sndMessage(text, m.channelID)

		case *slack.RTMError:
			fmt.Printf("[ERROR]: %s", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("[ERROR]: invalid authentication")
			break

		case *slack.DisconnectedEvent:
			if ev.Intentional {
				break
			}
		}
	}
}

func (m *slackMessenger) sndMessage(text string, channelID string) {
	m.rtm.SendMessage(m.rtm.NewOutgoingMessage(text, channelID))
}
