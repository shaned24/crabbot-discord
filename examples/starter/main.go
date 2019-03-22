package main

import (
	"flag"
	"github.com/shaned24/crabbot-discord/crabbot"
	"github.com/shaned24/crabbot-discord/examples/starter/myRoutes"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
var (
	Token  string
	Prefix string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.StringVar(&Prefix, "p", "!", "Bot Prefix")
	flag.Parse()
}

func main() {
	bot, err := crabbot.NewBot(
		Token,
		Prefix,
		myRoutes.NewSubRoute(),
		myRoutes.NewPing(),
		myRoutes.NewAvatar(),
		myRoutes.NewUser(),
	)

	if err != nil {
		log.Println("error creating Bot session,", err)
	}

	defer bot.Close()

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
