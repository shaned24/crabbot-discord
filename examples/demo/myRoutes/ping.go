package myRoutes

import (
	"github.com/Necroforger/dgrouter"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/shaned24/crabbot-discord/crabbot"
	"log"
)

const pingRoute = "ping"
const pingDescription = "responds with pong"

type Ping struct{}

func (p *Ping) Register(router *exrouter.Route) *exrouter.Route {
	return router.On(p.GetRouteCommand(), p.Handle)
}

func (p *Ping) GetSubRoutes() []crabbot.RouteHandler {
	return nil
}

func (p *Ping) Handle(ctx *exrouter.Context) {
	_, err := ctx.Reply("pong")
	if err != nil {
		log.Print("Something went wrong when handling Ping request", err)
	}
}

func (p *Ping) GetRouteCommand() string {
	return pingRoute
}

func (p *Ping) SetDescription(router *exrouter.Route) *dgrouter.Route {
	return router.Desc(pingDescription)
}

func NewPing() *Ping {
	return &Ping{}
}
