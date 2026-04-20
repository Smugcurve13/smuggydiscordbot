package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	fmt.Println("Message Received")
	msg := message.Content
	if strings.HasPrefix(msg, "!") {
		switch msg {
		case "!ping" :
			session.ChannelMessageSend(message.ChannelID, "This is ping function")
		case "!stats" :
			stats, err := getStats()
			if err != nil {
				session.ChannelMessageSend(message.ChannelID, "Error Getting Stats")
				return
			}
			session.ChannelMessageSend(message.ChannelID, stats)
		case "!help" :
			session.ChannelMessageSend(message.ChannelID, "This is help function")
		}
	}
}