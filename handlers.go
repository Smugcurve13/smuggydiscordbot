package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var WHITELISTED_IDS = []string{"786964843315986452" , "660947929057722388"}

var COMMAND_REGISTRY = map[string]func(*discordgo.MessageCreate, string)string{"help": helpFunc, "ping": pingFunc, "stats": statsFunc, "run":runFunc} 

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
// Deprecated: MessageHandlerv2 is inefficient and will not be updated.
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
			userID := message.Author.ID
			found := false
			for _, id := range WHITELISTED_IDS {
				if userID == id {
					found = true
					break
				}
			}
			if found {
				BLACKLIST_CMDS := []string{"rm -rf", "mkfs", "dd", "shutdown","test"}
				for _, cmd := range BLACKLIST_CMDS {
					if strings.Contains(argument, cmd) {
						session.ChannelMessageSend(message.ChannelID, "Galat command dalta hai! Try Again")
						return
					}
				}
				output, err := runLocalCommand(argument)
				if err != nil {
					fmt.Println("Error")
				}
				output = fmt.Sprintf("``` %s ```", output)
				session.ChannelMessageSend(message.ChannelID, output)
			} else {
				session.ChannelMessageSend(message.ChannelID, "Not Authorised")
			}
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

func MessageHandlerv3(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	fmt.Println("Message Received")
	msg := message.Content
	if strings.HasPrefix(msg, "!") {
		msg := strings.TrimPrefix(msg, "!")
		user_command, argument, _ := strings.Cut(msg, " ")
		if cmd_func, exists := COMMAND_REGISTRY[user_command]; exists {
			output := cmd_func(message, argument)
			session.ChannelMessageSend(message.ChannelID, output)
		}
	}
}