package handler

import (
	"fmt"

	"github.com/kobayashi-m42/go-slack-bot/domain/service"

	"github.com/nlopes/slack"
)

type Services struct {
	GitHubRepositoryService service.GitHubRepositoryService
}

type slackBot struct {
	botID        string
	channelID    string
	rtm          *slack.RTM
	messageEvent chan *slack.MessageEvent
	Services
}

type Bot interface {
	Listen()
	sndMessage(text string, channelID string)
}

func NewBot(token, botID, channelID string, services *Services) Bot {
	api := slack.New(token)
	rtm := api.NewRTM()

	return &slackBot{
		botID:        "<@" + botID + ">",
		channelID:    channelID,
		rtm:          rtm,
		messageEvent: make(chan *slack.MessageEvent),
		Services:     *services,
	}
}

func (s *slackBot) Listen() {
	go s.rtm.ManageConnection()

	go func() {
		for {
			ev := <-s.messageEvent
			go s.handleMessage(ev)
		}
	}()

	for msg := range s.rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			s.messageEvent <- ev

		case *slack.RTMError:
			fmt.Printf("[ERROR]: %s", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("[ERROR]: %s", authenticationError)
			break

		case *slack.DisconnectedEvent:
			if ev.Intentional {
				break
			}
		}
	}
}

func (s *slackBot) sndMessage(msg string, channelID string) {
	s.rtm.SendMessage(s.rtm.NewOutgoingMessage(msg, channelID))
}
