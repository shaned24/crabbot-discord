package myRoutes

import (
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/shaned24/crabbot-discord/crabbot/chat"
	"log"
	"time"
)

type RequestInput struct {
	Timeout time.Duration
}

func NewQueryInput() *RequestInput {
	return &RequestInput{
		Timeout: 10,
	}
}

func (r *RequestInput) Handle(ctx *exrouter.Context) {
	c, err := ctx.Ses.UserChannelCreate(ctx.Msg.Author.ID)
	if err != nil {
		log.Printf("Could not create direct channel to user: %v", err)
	}

	msg, err := chat.RequestInput(ctx, c.ID, "Are you there?", r.Timeout*time.Second)

	if err != nil {
		log.Printf("Something went wrong while requesting input: %v", err)

	}

	if msg != nil {
		_, _ = ctx.Ses.ChannelMessageSend(c.ID, fmt.Sprintf("You replied with: %s", msg.Content))
		if err != nil {
			log.Printf("Something went wrong sending message to channel: %v", err)
		}
	} else {
		_, _ = ctx.Ses.ChannelMessageSend(c.ID, "Timeout")
		if err != nil {
			log.Printf("Something went wrong sending message to channel: %v", err)
		}
	}

}

func (r *RequestInput) GetCommand() string {
	return "requestInput"
}

func (r *RequestInput) GetDescription() string {
	return fmt.Sprintf("Receive a direct message from the bot that "+
		"requests user input with a timeout of %d seconds", r.Timeout)
}
