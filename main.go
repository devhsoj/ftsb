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
		log.Printf("failed to load from .env: %s", err)
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

		log.Printf("stopped bot @ %s", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
	}()

	if err != nil {
		log.Fatalf("failed to create new Discord session: %s", err)
	}

	dg.AddHandler(MessageCreate)
	dg.Identify.Intents = discordgo.IntentGuildMessages

	if err = dg.Open(); err != nil {
		log.Fatalf("failed to open Discord session: %s", err)
	}

	log.Printf("started bot @ %s", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
