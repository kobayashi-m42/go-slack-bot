package main

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

type Bot struct {
	botID     string
	channelID string
}

func NewBot(botID, channelID string) *Bot {
	return &Bot{
		botID:     "<@" + botID + ">",
		channelID: channelID,
	}
}

// バリデーション
func (b *Bot) ValidateMessageEvent(ev *slack.MessageEvent) error {
	// channelが異なる場合エラー
	if ev.Channel != b.channelID {
		return fmt.Errorf("invalid channel %s %s", ev.Channel, ev.Msg.Text)
	}

	// bot宛て以外はエラー
	if !strings.HasPrefix(ev.Msg.Text, b.botID) {
		return fmt.Errorf("not a message for bot %s %s", ev.Channel, ev.Msg.Text)
	}

	// メッセージが空の場合エラー
	m := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
	if len(m) == 0 {
		return fmt.Errorf("invalid message")
	}
	return nil
}

func (b *Bot) GetTitleFromText(text string) string {
	removedBotID := strings.ReplaceAll(text, b.botID, "")
	message := b.sanitizeMsg(removedBotID)
	return message
}

func (b *Bot) sanitizeMsg(msg string) string {
	msg = strings.TrimSpace(msg)
	return msg
}
