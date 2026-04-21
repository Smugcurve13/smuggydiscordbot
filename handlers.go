package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var WHITELISTED_IDS = []string{"786964843315986452" , "660947929057722388"}

// Deprecated: MessageHandler is inefficient and will not be updated.
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
		if msg[:5] == "!run" {
			userID := message.Author.ID
			found := false
			for _, id := range WHITELISTED_IDS {
				if userID == id {
					found = true
					break
				}
			}
			if found {
				session.ChannelMessageSend(message.ChannelID, msg)

			} else {
				session.ChannelMessageSend(message.ChannelID, "Not Authorised")
			}
		}
		}
	}
}

func MessageHandlerv2(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	fmt.Println("Message Received")
	msg := message.Content
	if strings.HasPrefix(msg, "!") {
		msg := strings.TrimPrefix(msg, "!")
		command, argument, _ := strings.Cut(msg, " ")
		if command == "run" {
			fmt.Println(argument)
			output, err := runLocalCommand(argument)
			if err != nil {
				fmt.Println("Error")
			}
			output = fmt.Sprintf("``` %s ```", output)
			session.ChannelMessageSend(message.ChannelID, output)
		}
		if command == "stats" {
			stats, err := getStats()
			if err != nil {
				session.ChannelMessageSend(message.ChannelID, "Error Getting Stats")
				return
			}
			session.ChannelMessageSend(message.ChannelID, stats)
		}
		if command == "help" {
			session.ChannelMessageSend(message.ChannelID, "This is help function")
		}
		if command == "ping" {
			session.ChannelMessageSend(message.ChannelID, "This is ping function")
		}
	}
}