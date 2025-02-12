package discord

import (
	"log"
	"os"
	"time"

	"github.com/ElHefe3/AngloLex/pkg/wordnik"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

type DailyPoster struct {
	session   *discordgo.Session
	channelID string
}

func NewDailyPoster(session *discordgo.Session) *DailyPoster {
	channelID := os.Getenv("DISCORD_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("[NewDailyPoster] Missing DISCORD_CHANNEL_ID in environment variables.")
	}

	return &DailyPoster{
		session:   session,
		channelID: channelID,
	}
}

func (dp *DailyPoster) StartScheduler() {
	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(1).Day().At("10:00").Do(dp.PostWordOfTheDay)
	if err != nil {
		log.Fatalf("[StartScheduler] Error scheduling daily post: %v", err)
	}

	scheduler.StartAsync()
	log.Println("[StartScheduler] Daily posting job scheduled at 10:00 UTC.")
}

func (dp *DailyPoster) PostWordOfTheDay() {
	if dp.channelID == "" {
		log.Println("[PostWordOfTheDay] No channel ID specified, skipping daily post.")
		return
	}

	wordOfTheDay := wordnik.GetWordOfTheDay()

	_, err := dp.session.ChannelMessageSend(dp.channelID, wordOfTheDay)
	if err != nil {
		log.Printf("[PostWordOfTheDay] Failed to send message: %v", err)
	} else {
		log.Println("[PostWordOfTheDay] Successfully posted Word of the Day.")
	}
}
