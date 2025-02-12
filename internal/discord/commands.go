package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RegisterCommands() {
	if b.session == nil {
		log.Fatal("Bot session is nil. Ensure session is initialized before registering commands.")
	}

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "word-of-the-day",
			Description: "Get a random word of the day from Wikipedia.",
			Type:        discordgo.ChatApplicationCommand,
		},
	}

	for _, cmd := range commands {
		_, err := b.session.ApplicationCommandCreate(b.session.State.User.ID, b.guildID, cmd)
		if err != nil {
			log.Printf("Error registering command %s: %v", cmd.Name, err)
		} else {
			log.Printf("Command %s registered", cmd.Name)
		}
	}
}
