package myRoutes

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/shaned24/crabbot-discord/crabbot"
	"log"
)

type ExampleSubRoute struct{}

func (u *ExampleSubRoute) GetSubRoutes() []crabbot.Route {
	return []crabbot.Route{
		NewUser(),
	}
}

func (u *ExampleSubRoute) GetDescription() string {
	return "this is an example of using sub routes"
}

func (u *ExampleSubRoute) Handle(ctx *exrouter.Context) {
	_, err := ctx.Reply("This is a sub route." + ctx.Msg.Author.Username)
	if err != nil {
		log.Printf("Something went wrong: %v", err)
	}
}

func (u *ExampleSubRoute) GetCommand() string {
	return "sub"
}

func NewSubRoute() *ExampleSubRoute {
	return &ExampleSubRoute{}
}
