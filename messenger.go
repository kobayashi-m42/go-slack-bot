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
	bot Bot
}

func NewMessenger(token string, bot *Bot) Messenger {
	api := slack.New(token)
	rtm := api.NewRTM()

	return &slackMessenger{
		rtm: rtm,
		bot: *bot,
	}
}

func (m *slackMessenger) Listen() {
	go m.rtm.ManageConnection()

	for msg := range m.rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := m.bot.ValidateMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle responseMessage: %s", err)
				break
			}
			text, channelID := m.bot.createResponseMessage(ev)
			m.sndMessage(text, channelID)

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
