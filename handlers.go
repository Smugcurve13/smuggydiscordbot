package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var COMMAND_REGISTRY = map[string]func(*discordgo.MessageCreate, string)string{"help": helpFunc, "ping": pingFunc, "stats": statsFunc, "run":runFunc} 

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
			output = fmt.Sprintf("```%s```",output)
			session.ChannelMessageSend(message.ChannelID, output)
		}
	}
}