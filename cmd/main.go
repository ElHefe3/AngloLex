package main

import (
	"log"
	"os"

	"github.com/ElHefe3/AngloLex/internal/discord"
	"github.com/joho/godotenv"
)

const (
	DISCORD_TOKEN_KEY   = "DISCORD_TOKEN"
	GUILD_ID_KEY        = "GUILD_ID"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[main] Error loading .env file")
	}

	discordToken := os.Getenv(DISCORD_TOKEN_KEY)
	if discordToken == "" {
		log.Fatalf("[main] Discord bot token not found in environment variables")
	}

	guildId := os.Getenv(GUILD_ID_KEY)
	if guildId == "" {
		log.Fatalf("[main] Guild ID not found in environment variables")
	}

	bot, err := discord.NewBot(discordToken, guildId)
	if err != nil {
		log.Fatalf("[main] Failed to initialize bot: %v", err)
	}

	bot.RegisterCommands()
	bot.RegisterHandlers()

	dailyPoster := discord.NewDailyPoster(bot.GetSession())
	dailyPoster.StartScheduler()

	log.Println("[main] Bot is running. Press Ctrl+C to stop.")
	select {}
}
