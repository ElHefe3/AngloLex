package discord

import (
	"fmt"
	"log"

	"github.com/ElHefe3/AngloLex/pkg/wordnik"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) RegisterHandlers() {
	if b.session == nil {
		log.Fatal("[RegisterHandlers] Bot session is nil. Ensure session is initialized before registering handlers.")
	}

	b.session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("[RegisterHandlers] Bot is now running as", s.State.User.Username)
	})

	b.session.AddHandler(b.handleInteractionCreate)
}

func (b *Bot) handleInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("[handleInteractionCreate] Received an interaction!")
	log.Printf("[handleInteractionCreate] Interaction Type: %v", i.Type)

	if i.Type != discordgo.InteractionApplicationCommand {
		log.Println("[handleInteractionCreate] Skipping non-command interaction.")
		return
	}

	commandName := i.ApplicationCommandData().Name
	log.Printf("[handleInteractionCreate] Command Name: %s", commandName)

	switch commandName {
	case "word-of-the-day":
		b.handleWordOfTheDay(s, i)
	default:
		log.Printf("[handleInteractionCreate] Unknown command: %s", commandName)
	}
}

func (b *Bot) handleWordOfTheDay(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("[handleWordOfTheDay] Received /word-of-the-day command")

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		log.Printf("[handleWordOfTheDay] Failed to acknowledge interaction: %v", err)
		return
	}

	todaysWord := wordnik.GetWordOfTheDay()
	if todaysWord == "" {
		log.Println("[handleWordOfTheDay] Wordnik API returned an empty response.")
		todaysWord = "⚠️ Could not fetch the word of the day. Please try again later."
	}

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &todaysWord,
	})
	if err != nil {
		log.Printf("[handleWordOfTheDay] Failed to send final response: %v", err)

		_, sendErr := s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Content: fmt.Sprintf("📖 Here's your word of the day:\n%s", todaysWord),
		})
		if sendErr != nil {
			log.Printf("[handleWordOfTheDay] Failed to send follow-up message: %v", sendErr)
		}
	} else {
		log.Println("[handleWordOfTheDay] Response successfully sent!")
	}
}
