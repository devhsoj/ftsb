package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		if os.Getenv("DOCKER") != "true" {
			log.Printf("failed to load from .env: %s", err)
		}
	}

	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")

	if len(discordBotToken) == 0 {
		log.Fatal("DISCORD_BOT_TOKEN is not defined in the environment (and .env) or is invalid!")
	}

	dg, err := discordgo.New("Bot " + discordBotToken)

	defer func() {
		if err := dg.Close(); err != nil {
			log.Fatalf("failed to properly close Discord session: %s", err)
		}

		log.Println("[+] stopped ftsb bot")
	}()

	if err != nil {
		log.Fatalf("failed to create new Discord session: %s", err)
	}

	dg.AddHandler(MessageCreate)

	dg.Identify.Intents |= discordgo.IntentGuilds
	dg.Identify.Intents |= discordgo.IntentGuildMessages

	if err = dg.Open(); err != nil {
		log.Fatalf("failed to open Discord session: %s", err)
	}

	log.Println("[+] started ftsb bot")

	go func() {
		for {
			time.Sleep(time.Hour)

			dg.State.Lock()

			for _, guild := range dg.State.Guilds {
				for _, channel := range guild.Channels {
					if channel.Type == discordgo.ChannelTypeGuildText && channel.Name == "status" {
						summary, err := GetTrailStatusSummary()

						if err != nil {
							log.Printf("failed to get trail status summary: %s", err)
						}

						if _, err = dg.ChannelMessageSend(channel.ID, summary); err != nil {
							log.Printf("failed to send trail status summary to channel '%s': %s", channel.ID, err)
						}

						log.Printf("sent updated summary to %s:%s", guild.Name, channel.Name)
					}
				}
			}

			dg.State.Unlock()
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
