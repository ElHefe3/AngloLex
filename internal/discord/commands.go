package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RegisterCommands() {
	if b.session == nil {
		log.Fatal("[RegisterCommands] Bot session is nil. Ensure session is initialized before registering commands.")
	}

	commands, err := b.session.ApplicationCommands(b.session.State.User.ID, b.guildID)
	if err != nil {
		log.Printf("[RegisterCommands] Failed to fetch existing commands: %v", err)
	} else {
		for _, cmd := range commands {
			err := b.session.ApplicationCommandDelete(b.session.State.User.ID, b.guildID, cmd.ID)
			if err != nil {
				log.Printf("[RegisterCommands] Failed to delete command %s: %v", cmd.Name, err)
			} else {
				log.Printf("[RegisterCommands] Deleted old command: %s", cmd.Name)
			}
		}
	}

	newCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "word-of-the-day",
			Description: "Get a random word of the day from Wikipedia.",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "define",
			Description: "Look up a word's definition.",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "word",
					Description: "The word to define",
					Required:    true,
				},
			},
		},
		{
			Name:			"find-word-for",
			Description:    "Find a word that fits a given sentence.",
            Type:            discordgo.ChatApplicationCommand,
            Options:         []*discordgo.ApplicationCommandOption{
				{
					Type: 		 discordgo.ApplicationCommandOptionString,
					Name:        "sentence",
                    Description: "The sentence to find a word for",
                    Required:    true,
				},
			},
		},
	}

	for _, cmd := range newCommands {
		_, err := b.session.ApplicationCommandCreate(b.session.State.User.ID, b.guildID, cmd)
		if err != nil {
			log.Printf("[RegisterCommands] Error registering command %s: %v", cmd.Name, err)
		} else {
			log.Printf("[RegisterCommands] Command %s registered", cmd.Name)
		}
	}
}
