package crabbot

import (
	"github.com/Necroforger/dgrouter/exrouter"
)

// RouteHandler is an interface that wraps the exrouter.Route implementation
//
// The register function registers a route for a router using the output of the GetRouteCommand function to specify
// the anchor that will trigger the Handle function when invoked with the string returned from GetRouteCommand
//
// The output of the GetDescription command should return a short description for that this handler does
// this description will be used when registering route handlers with the exrouter.Route implementation
type RouteHandler interface {
	Register(router *exrouter.Route) *exrouter.Route
	Handle(ctx *exrouter.Context)
	GetRouteCommand() string
	GetDescription() string
}

// SubRouteHandler is a small extension to the RouterHandler
//
// The GetSubRoutes function will return an array of RouterHandlers which will be used to attach these RouteHandlers
// as children to the parent exrouter.Route
type SubRouteHandler interface {
	RouteHandler
	GetSubRoutes() []RouteHandler
}

// Registers a help handler for the passed in router
func registerHelpHandler(router *exrouter.Route) {
	handler := NewHelp()
	helpRoute := handler.Register(router)
	helpRoute.Desc(handler.GetDescription())
}

// Attaches a description to the router that will be used for printing help descriptions and usages
func registerRouterWithDescription(handler RouteHandler, router *exrouter.Route) *exrouter.Route {
	route := handler.Register(router)
	route.Desc(handler.GetDescription())
	return route
}

// Registers all RouteHandlers and SubRouteHandlers to the main router passed in
// This function adds a default help handler for all routes and sub routes
// Since the RouteHandler interface requires a GetDescription function to be implemented
// these descriptions will be added to each help handler respectively
func RegisterRoutes(router *exrouter.Route, routeHandlers ...RouteHandler) {
	for _, routeHandler := range routeHandlers {
		// Register root help
		registerHelpHandler(router)

		// Register root route
		rootRouter := registerRouterWithDescription(routeHandler, router)

		// check if the root route as sub routes
		if subRouteHandler, ok := routeHandler.(SubRouteHandler); ok {
			if subRouteHandlers := subRouteHandler.GetSubRoutes(); subRouteHandlers != nil {
				// register help command for sub route
				registerHelpHandler(rootRouter)

				for _, handler := range subRouteHandlers {
					// Register the sub route onto the root route
					registerRouterWithDescription(handler, rootRouter)
				}
			}
		}
	}
}
