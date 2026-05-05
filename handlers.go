package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var COMMAND_REGISTRY = map[string]func(*discordgo.Session, *discordgo.MessageCreate, string)string{"ai": testaiFunc,"help": helpFunc, "ping": pingFunc, "stats": statsFunc, "run":runFunc, "roast":roastFuncv2, "quiz": quizFunc } 

var SERVER_ALLOWED_COMMANDS = map[string][]string{
	"1495404097535479958": {"ai", "help", "ping", "stats", "run", "roast", "quiz"}, // SmuggyDen
	"1340357023820415048": {"roast", "ping"}, // BadTrip Server
	"786879993975144458" : {"roast", "ping"}, // Vibes Server (testing ground for command bifurcation)
}

func MessageHandlerv3(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	fmt.Println("Message Received")
	msg := message.Content
	userCommand, argument, isCommand := messageParser(msg)
	// check if command is allowed in that guild
	isAllowed := ifGuildAllowed(message.GuildID, userCommand)
	if isCommand && isAllowed {
		if cmd_func, exists := COMMAND_REGISTRY[userCommand]; exists {
			output := cmd_func(session, message, argument)
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

func ifGuildAllowed(guildID string, userCommand string) bool {
	allowedCommands := SERVER_ALLOWED_COMMANDS[guildID]
	isAllowed := false
	for _, command := range allowedCommands {
		if userCommand == command {
		isAllowed = true
		break
		}
	}
	return isAllowed
}