package crabbot

import (
	"github.com/Necroforger/dgrouter/exrouter"
)

// Route is an interface that wraps the exrouter.Route implementation
//
// The register function registers a route for a router using the output of the GetRouteCommand function to specify
// the anchor that will trigger the Handle function when invoked with the string returned from GetRouteCommand
//
// The output of the GetDescription command should return a short description for that this handler does
// this description will be used when registering route handlers with the exrouter.Route implementation
type Route interface {
	Register(router *exrouter.Route) *exrouter.Route
	Handle(ctx *exrouter.Context)
	GetRouteCommand() string
	GetDescription() string
}

// SubRoute is a small extension to the Route
//
// The GetSubRoutes function will return an array of Routes which will be used to attach these Routes
// as children to the parent exrouter.Route
type SubRoute interface {
	Route
	GetSubRoutes() []Route
}

// registerHelpHandler Registers a help handler for the passed in router
func registerHelpHandler(router *exrouter.Route) {
	handler := NewHelp()
	helpRoute := handler.Register(router)
	helpRoute.Desc(handler.GetDescription())
}

// registerRouterWithDescription Attaches a description to the router that will be used for printing help descriptions and usages
func registerRouterWithDescription(handler Route, router *exrouter.Route) *exrouter.Route {
	route := handler.Register(router)
	route.Desc(handler.GetDescription())
	return route
}

// RegisterRoutes Registers all Routes and SubRoutes to the main router passed in
// This function adds a default help handler for all routes and sub routes
// Since the Route interface requires a GetDescription function to be implemented
// these descriptions will be added to each help handler respectively
func RegisterRoutes(router *exrouter.Route, routes ...Route) {
	for _, route := range routes {
		// Register root help
		registerHelpHandler(router)

		// Register root route
		rootRouter := registerRouterWithDescription(route, router)

		// check if the root route as sub routes
		if subRoute, ok := route.(SubRoute); ok {
			if subRouteHandlers := subRoute.GetSubRoutes(); subRouteHandlers != nil {
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
