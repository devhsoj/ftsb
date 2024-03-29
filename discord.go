package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func MessageCreate(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == sess.State.User.ID {
		return
	}

	if msg.Content == "!trailstatus" {
		statusSummary, err := GetTrailStatusSummary()

		if err != nil {
			log.Printf("failed to get trail status summary: %s", err)
			return
		}

		if _, err := sess.ChannelMessageSend(msg.ChannelID, statusSummary); err != nil {
			log.Printf("failed to send trail status summary to channel '%s': %s", msg.ChannelID, err)
			return
		}
	}
}
