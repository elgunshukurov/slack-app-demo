package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"log"
	"os"
	"strings"
)

func main() {
	appToken := os.Getenv("SLACK_APP_TOKEN")
	if appToken == "" {
		panic("SLACK_APP_TOKEN must be set.\n")
	}

	if !strings.HasPrefix(appToken, "xapp-") {
		panic("SLACK_APP_TOKEN must have the prefix \"xapp-\".")
	}

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		panic("SLACK_BOT_TOKEN must be set.\n")
	}

	if !strings.HasPrefix(botToken, "xoxb-") {
		panic("SLACK_BOT_TOKEN must have the prefix \"xoxb-\".")
	}

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	confirmationBlock := slack.NewActionBlock(
		"",
		slack.NewButtonBlockElement(
			"",
			"yes_button",
			slack.NewTextBlockObject("plain_text", "Yes", false, false),
		).WithStyle(slack.StylePrimary),
		slack.NewButtonBlockElement(
			"",
			"no_button",
			slack.NewTextBlockObject("plain_text", "No", false, false),
		).WithStyle(slack.StyleDanger),
	)

	headerText := slack.NewTextBlockObject("mrkdwn", "*Confirmation Needed*", false, false)
	headerSelection := slack.NewSectionBlock(headerText, nil, nil)

	message := slack.MsgOptionBlocks(
		headerSelection,
		slack.NewDividerBlock(),
		slack.NewSectionBlock(
			slack.NewTextBlockObject("mrkdwn", "Please confirm your action", false, false),
			nil,
			nil,
		),
		slack.NewDividerBlock(),
		confirmationBlock,
	)

	api.SendMessage("#demo", message)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	socketmodeHandler := socketmode.NewSocketmodeHandler(client)
	socketmodeHandler.Handle(socketmode.EventTypeInteractive, middlewareInteractive)

	socketmodeHandler.RunEventLoop()
}

func middlewareInteractive(evt *socketmode.Event, client *socketmode.Client) {
	callback, ok := evt.Data.(slack.InteractionCallback)
	if !ok {
		fmt.Printf("Ignored %+v\n", evt)
		return
	}

	fmt.Printf("Interaction received: %+v\n", callback)

	var payload interface{}

	switch callback.Type {
	case slack.InteractionTypeBlockActions:
		// See https://api.slack.com/apis/connections/socket-implement#button
		client.Debugf("button clicked!")
	case slack.InteractionTypeShortcut:
	case slack.InteractionTypeViewSubmission:
		// See https://api.slack.com/apis/connections/socket-implement#modal
	case slack.InteractionTypeDialogSubmission:
	default:

	}

	client.Ack(*evt.Request, payload)
}
