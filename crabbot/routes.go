package crabbot

import (
	"github.com/Necroforger/dgrouter/exrouter"
)

// Route is an interface that wraps the exrouter.Route implementation
//
// The register function registers a route for a router using the output of the GetCommand function to specify
// the anchor that will trigger the Handle function when invoked with the string returned from GetCommand
//
// The output of the GetDescription command should return a short description for that this handler does
// this description will be used when registering route handlers with the exrouter.Route implementation
type Route interface {
	// GetCommand return the command that will be used to hook into the router
	GetCommand() string
	// GetDescription return a description of your route
	GetDescription() string
}

type DefaultRouteHandler interface {
	// Handle The call back that will be ran when the user inputted command matches a route definition
	Handle(ctx *exrouter.Context)
}

type CustomRouteHandler interface {
	// Register a means to add more in depth routing definitions such as regex matching
	Register(router *exrouter.Route) *exrouter.Route
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
	route := NewHelp(router)
	registerRoute(route, router)
}

// registerRoute Attaches a description to the router that will be used for printing help descriptions and usages
func registerRoute(route Route, router *exrouter.Route) *exrouter.Route {
	prefix := route.GetCommand()

	var newRouter *exrouter.Route

	if handlerHandler, ok := route.(DefaultRouteHandler); ok {
		newRouter = router.On(prefix, handlerHandler.Handle)
	}
	if routeHandler, ok := route.(CustomRouteHandler); ok {
		newRouter = routeHandler.Register(router)
	}
	newRouter.Desc(route.GetDescription())
	return newRouter
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
		rootRouter := registerRoute(route, router)

		// check if the root route as sub routes
		if subRoute, ok := route.(SubRoute); ok {
			if subRouteHandlers := subRoute.GetSubRoutes(); subRouteHandlers != nil {
				// register help command for sub route
				registerHelpHandler(rootRouter)

				for _, handler := range subRouteHandlers {
					// Register the sub route onto the root route
					registerRoute(handler, rootRouter)
				}
			}
		}
	}
}
