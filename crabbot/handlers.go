package crabbot

import "github.com/Necroforger/dgrouter/exrouter"

type RouteHandler interface {
	Register(router *exrouter.Route)
	Handle(ctx *exrouter.Context)
	GetRouteCommand() string
	GetDescription() string
}

func RegisterRoutes(router *exrouter.Route, routeHandlers ...RouteHandler) {
	for _, r := range routeHandlers {
		r.Register(router)
	}
}