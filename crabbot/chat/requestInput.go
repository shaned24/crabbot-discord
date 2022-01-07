package chat

import (
	"errors"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"time"
)

func RequestInput(ctx *exrouter.Context, channelID string, promptMessage string, timeout time.Duration) (*discordgo.Message, error) {
	msg, err := ctx.Ses.ChannelMessageSend(channelID, promptMessage)
	if err != nil {
		return nil, err
	}

	// delete prompted message at the end
	defer func() {
		_ = ctx.Ses.ChannelMessageDelete(msg.ChannelID, msg.ID)
	}()

	// Channel that we will timeout on if no input is received
	timeoutChan := make(chan int)
	go func() {
		time.Sleep(timeout)
		timeoutChan <- 0
	}()

	for {
		select {
		// run if the input is received
		case userMessage := <-NewMessageCreateChan(ctx.Ses):
			if userMessage.Author.ID != ctx.Msg.Author.ID {
				continue
			}
			_ = ctx.Ses.ChannelMessageDelete(userMessage.ChannelID, userMessage.ID)
			return userMessage.Message, nil
		// wait for timeout to exceed limit
		case <-timeoutChan:
			return nil, errors.New("timeout on requesting input")
		}
	}
}

func NewMessageCreateChan(session *discordgo.Session) chan *discordgo.MessageCreate {
	out := make(chan *discordgo.MessageCreate)
	session.AddHandlerOnce(func(_ *discordgo.Session, event *discordgo.MessageCreate) {
		out <- event
	})
	return out
}
