package myRoutes

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/shaned24/crabbot-discord/crabbot"
	"log"
)

type SubRoute struct{}

func (u *SubRoute) Register(router *exrouter.Route) *exrouter.Route {
	return router.On(u.GetRouteCommand(), u.Handle)
}

func (u *SubRoute) GetSubRoutes() []crabbot.RouteHandler {
	return []crabbot.RouteHandler{
		NewUser(),
	}
}

func (u *SubRoute) GetDescription() string {
	return "this is an example of using sub routes"
}

func (u *SubRoute) Handle(ctx *exrouter.Context) {
	_, err := ctx.Reply("This is a sub route." + ctx.Msg.Author.Username)
	if err != nil {
		log.Printf("Something went wrong: %v", err)
	}
}

func (u *SubRoute) GetRouteCommand() string {
	return "sub"
}

func NewSubRoute() *SubRoute {
	return &SubRoute{}
}
