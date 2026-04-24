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
	user_command, argument, is_command := messageParser(msg)
	if is_command {
		if cmd_func, exists := COMMAND_REGISTRY[user_command]; exists {
			output := cmd_func(message, argument)
			output = fmt.Sprintf("```%s```",output)
			session.ChannelMessageSend(message.ChannelID, output)
		}
	}
}

func messageParser(msg string) (string, string, bool) {
	is_command := false
	if strings.HasPrefix(msg, "!") {
		msg := strings.TrimPrefix(msg, "!")
		user_command, argument, _ := strings.Cut(msg, " ")
		is_command = true
		return user_command, argument, is_command
	} else {
		return "", "", is_command
	}
}