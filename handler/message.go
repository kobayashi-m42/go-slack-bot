package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kobayashi-m42/go-slack-bot/domain"

	"github.com/nlopes/slack"
)

func (s *slackBot) handleMessage(ev *slack.MessageEvent) {
	if err := s.validateMessageEvent(ev); err != nil {
		s.log.Warnw(err.Error(),
			"channel", ev.Channel,
			"request", ev.Msg.Text,
		)
		return
	}

	title := s.getTitleFromText(ev.Msg.Text)
	botResponse := s.GitHubRepositoryService.CreateIssue(title)
	if err := botResponse.Err; err != nil {
		msg := ""

		switch e := err.(type) {
		case *domain.CreateIssueError:
			msg = fmt.Sprint(err, e.Msg)
		default:
			msg = fmt.Sprint(err)
		}

		s.sndMessage(msg, s.channelID)
	}

	s.sndMessage(botResponse.Msg, s.channelID)
}

func (s *slackBot) validateMessageEvent(ev *slack.MessageEvent) error {
	// 異なるチャンネル
	if ev.Channel != s.channelID {
		return errors.New(invalidChannelError)
	}

	// bot宛て以外
	if !strings.HasPrefix(ev.Msg.Text, s.botID) {
		return errors.New(invalidBotError)
	}

	// メッセージが空
	m := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
	if len(m) == 0 {
		return errors.New(messageEmptyError)
	}
	return nil
}

func (s *slackBot) getTitleFromText(text string) string {
	removedBotID := strings.ReplaceAll(text, s.botID, "")
	message := s.sanitizeMsg(removedBotID)
	return message
}

func (s *slackBot) sanitizeMsg(msg string) string {
	msg = strings.TrimSpace(msg)
	return msg
}
