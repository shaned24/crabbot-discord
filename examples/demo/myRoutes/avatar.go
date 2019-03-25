package myRoutes

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"log"
)

const avatarRoute = "avatar"
const avatarDescription = "returns the user's avatar"

type Avatar struct{}

func (a *Avatar) GetDescription() string {
	return avatarDescription
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

func NewAvatar() *Avatar {
	return &Avatar{}
}
