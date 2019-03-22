package crabbot

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Token   string
	Prefix  string
	Routes  []Route
	Session *discordgo.Session
}

func (b *Bot) registerRoutes(dg *discordgo.Session, Routes ...Route) {
	router := exrouter.New()

	RegisterRoutes(
		router,
		Routes...,
	)

	// Add message handler
	dg.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		_ = router.FindAndExecute(b.Session, b.Prefix, dg.State.User.ID, m.Message)
	})
}

func NewBot(token string, prefix string, Routes ...Route) (*Bot, error) {
	dgSession, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		Token:   token,
		Prefix:  prefix,
		Routes:  Routes,
		Session: dgSession,
	}

	bot.registerRoutes(bot.Session, Routes...)

	return bot, nil
}

func (b *Bot) Start() error {
	err := b.Session.Open()
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) Close() error {
	return b.Session.Close()
}
