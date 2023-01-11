package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {
	tkn := os.Getenv("SLACK_TOKEN")

	if tkn == "" {
		panic("slack token is not set")
	}

	channel := os.Getenv("SLACK_CHANNEL")

	if channel == "" {
		panic("slack channel is not set")
	}

	client := slack.New(tkn)

	msg := "Hello, Go!"

	if _, _, err := client.PostMessage(channel, slack.MsgOptionText(msg, true)); err != nil {
		panic(err)
	}
}
