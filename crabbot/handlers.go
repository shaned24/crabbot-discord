package crabbot

import (
	"github.com/Necroforger/dgrouter"
	"github.com/Necroforger/dgrouter/exrouter"
)

type RouteHandler interface {
	Register(router *exrouter.Route) *exrouter.Route
	Handle(ctx *exrouter.Context)
	GetRouteCommand() string
	GetSubRoutes() []RouteHandler
	SetDescription(router *exrouter.Route) *dgrouter.Route
}

func RegisterRoutes(router *exrouter.Route, routeHandlers ...RouteHandler) {
	for _, r := range routeHandlers {

		// Register root route
		rootRoute := r.Register(router)

		// Register root help
		rootHelp := NewHelp()
		rootHelpRoute := rootHelp.Register(router)
		rootHelp.SetDescription(rootHelpRoute)

		// check if the root route as sub routes
		subRoutes := r.GetSubRoutes()
		if subRoutes != nil {
			for _, sub := range subRoutes {
				// Register the sub route onto the root route
				subRoute := sub.Register(rootRoute)
				sub.SetDescription(subRoute)
			}

			// register help command for sub route
			subHelp := NewHelp()
			subHelpRoute := subHelp.Register(rootRoute)
			subHelp.SetDescription(subHelpRoute)

		}
		r.SetDescription(rootRoute)
	}
}
