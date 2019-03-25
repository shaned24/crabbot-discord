package myRoutes

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"log"
)

type DirectMessage struct{}

func NewDirectMessage() *DirectMessage {
	return &DirectMessage{}
}

func (d *DirectMessage) Register(router *exrouter.Route) *exrouter.Route {
	return router.On(d.GetRouteCommand(), d.Handle)
}

func (d *DirectMessage) Handle(ctx *exrouter.Context) {
	c, err := ctx.Ses.UserChannelCreate(ctx.Msg.Author.ID)
	if err != nil {
		log.Printf("Could not create direct channel to user: %v", err)
	}

	_, err = ctx.Ses.ChannelMessageSend(c.ID, "wow")
	if err != nil {
		log.Printf("Could not send message: %v", err)
	}
}

func (d *DirectMessage) GetRouteCommand() string {
	return "dm"
}

func (d *DirectMessage) GetDescription() string {
	return "Recieve a direct message from the bot"
}
