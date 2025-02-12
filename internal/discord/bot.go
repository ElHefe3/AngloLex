package discord

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
	guildID string
}

func NewBot(token, guildID string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	bot := &Bot{
		session: dg,
		guildID: guildID,
	}

	if err := bot.session.Open(); err != nil {
		return nil, errors.New("failed to open Discord session: " + err.Error())
	}

	return bot, nil
}

func (b *Bot) Start() error {
	log.Println("Bot is now running. Press Ctrl+C to exit.")
	return nil
}

func (b *Bot) GetSession() *discordgo.Session {
	return b.session
}

func (b *Bot) GetGuildID() string {
	return b.guildID
}
