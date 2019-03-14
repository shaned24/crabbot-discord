# Crab Bot Dicord

This library adds an interface to easily map routes to handlers
in a clean way.

This library depends on two libraries

- https://github.com/bwmarrin/discordgo
- https://github.com/Necroforger/dgrouter

## Installation

`TODO`

## Usage

Implementing a route is simple 

#### Implement the `crabbot.RouteHandler` interface:

**myroute.go**
```go

// Interface declaration
type RouteHandler interface {
	Register(router *exrouter.Route)
	Handle(ctx *exrouter.Context)
	GetRouteCommand() string
	GetDescription() string
}

...

// Implementation of the RouteHandler interface
type MyRoute struct {
	router *exrouter.Route
}

func (h *MyRoute) Handle(ctx *exrouter.Context) {
	// implement handler
}

func (h *MyRoute) GetRouteCommand() string {
	// We need to return the string that will trigger our handle function
	return "myroute"
}

func (h *MyRoute) GetDescription() string {
	// We can add a description 
	return "Does something after calling myroute"
}

func (h *MyRoute) Register(router *exrouter.Route) {
	router.On(h.GetRouteCommand(), h.Handle).Desc(h.GetDescription())
}

func NewMyRoute() *MyRoute {
	return &MyRoute{}
}
```

#### Add our implementation to a `[]RouteHandler`

Add the implementation to an array of `[]RouteHandler` and pass the array into the `crabbot.NewBot` function

**main.go**
```go
...

func main() {
    token := "my-bot-token"
    prefix := "!"
	
    // Create our routes
    handlers := []crabbot.RouteHandler{
        routes.MyRoute(),
    }
    
    bot, err := crabbot.NewBot(token, prefix, handlers...)
    ...
}
```


## Examples

- Examples can be found in the [./examples](https://github.com/shaned24/crabbot-discord/tree/master/examples/) directory
```go
package main

import (
	"flag"
	"github.com/shaned24/crabbot-discord/crabbot"
	"github.com/shaned24/crabbot-discord/crabbot/routes"
	"github.com/shaned24/crabbot-discord/examples/starter/myRoutes"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
var (
	Token string
	Prefix string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.StringVar(&Prefix, "p", "!", "Bot Prefix")
	flag.Parse()
}


func main() {
	// Create our routes
	handlers := []crabbot.RouteHandler{
		routes.NewHelp(),
	}
    
	
	// Create our bot session
	bot, err := crabbot.NewBot(Token, Prefix, handlers...)
	if err != nil {
		log.Println("error creating Bot session,", err)
	}

	defer bot.Close()

    // Start the bot
	err = bot.Start()
	if err != nil {
		log.Printf("Couldn't start the bot: %v", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
```