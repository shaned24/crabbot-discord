package myRoutes

import (
	"github.com/Necroforger/dgrouter"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/shaned24/crabbot-discord/crabbot"
	"log"
)

const avatarRoute = "avatar"
const avatarDescription = "returns the user's avatar"

type Avatar struct{}

func (a *Avatar) SetDescription(router *exrouter.Route) *dgrouter.Route {
	return router.Desc(avatarDescription)
}

func (a *Avatar) Register(router *exrouter.Route) *exrouter.Route {
	return router.On(avatarRoute, a.Handle)
}

func (a *Avatar) Handle(ctx *exrouter.Context) {
	_, err := ctx.Reply(ctx.Msg.Author.AvatarURL("2048"))
	if err != nil {
		log.Print("Something went wrong when handling Avatar request", err)
	}
}

func (a *Avatar) GetRouteCommand() string {
	return avatarRoute
}

func (a *Avatar) GetSubRoutes() []crabbot.RouteHandler {
	return nil
}

func NewAvatar() *Avatar {
	return &Avatar{}
}
