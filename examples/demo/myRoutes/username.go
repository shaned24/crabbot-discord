package myRoutes

import (
	"github.com/Necroforger/dgrouter"
	"github.com/Necroforger/dgrouter/exrouter"
	"log"
)

type User struct{}

func (u *User) Register(router *exrouter.Route) *exrouter.Route {
	return router.OnMatch("username", dgrouter.NewRegexMatcher("user(name)?"), u.Handle)
}

func (u *User) GetDescription() string {
	return "returns the users username"
}

func (u *User) Handle(ctx *exrouter.Context) {
	_, err := ctx.Reply("Your username is " + ctx.Msg.Author.Username)
	if err != nil {
		log.Printf("Something went wrong: %v", err)
	}
}

func (u *User) GetRouteCommand() string {
	return "user(name)"
}

func NewUser() *User {
	return &User{}
}
