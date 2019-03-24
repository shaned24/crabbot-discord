# Crab Bot Dicord

This library adds an interface to easily map routes to handlers
in a clean way.

This library depends on two libraries

- https://github.com/bwmarrin/discordgo
- https://github.com/Necroforger/dgrouter

## Installation

First things first `go get github.com/shaned24/crabbot-discord/crabbot`

## Usage

In a discord chat lets say you want your bot to respond to a `!ping` that was typed in a chatroom

#### Implement the `crabbot.Route` interface:
```go
// Interface declaration
type Route interface {
	Register(router *exrouter.Route) *exrouter.Route
	Handle(ctx *exrouter.Context)
	GetRouteCommand() string
	GetDescription(router *exrouter.Route) string
}
```

#### Sample implementation
```go
type PingRoute struct{}

func (u *PingRoute) Register(router *exrouter.Route) *exrouter.Route {
	return router.On(u.GetRouteCommand(), u.Handle)
}

func (u *PingRoute) GetDescription(router *exrouter.Route) string {
	return "This will respond with Pong!"
}

func (u *PingRoute) Handle(ctx *exrouter.Context) {
	ctx.Reply("Pong!")
}

func (u *PingRoute) GetRouteCommand() string {
	return "ping"
}

func NewPingRoute() *PingRoute {
	return &PingRoute{}
}
```

#### Example

```go
func main() {
    token := "my-bot-token"
    prefix := "!"
    
    // Create the bot instance
    bot, err := crabbot.NewBot(Token, Prefix, NewPingRoute())
    if err != nil {
        log.Println("error creating Bot session,", err)
    }

    defer bot.Close()
    err = bot.Start()
    
    // Wait here until CTRL-C or other term signal is received.
    log.Println("Bot is now running. Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc
}
```

## Examples
More examples can be found in the [examples](https://github.com/shaned24/crabbot-discord/tree/master/examples/) directory
