package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func helpFunc(message *discordgo.MessageCreate, arg string) string {
	return "Available commands: !ping, !help, !stats and !run <command> <argument>"
}

func pingFunc(message *discordgo.MessageCreate, arg string) string {
	return `Pong ! Type "!run" to get all available commands`
}

func statsFunc(message *discordgo.MessageCreate, arg string) string {
	stats, err := getStats()
	if err != nil {
		return "Unable to fetch system stats"
	}
	return stats
}

func runFunc(message *discordgo.MessageCreate, argument string) string {
	userID := message.Author.ID
	found := false
	WHITELISTED_IDS := getWhitelistedIDS()
	for _, id := range WHITELISTED_IDS {
		if userID == id {
			found = true
			break
		}
	}
	if found {
		argument = strings.TrimSpace(argument)
		if argument == "" {
			msg := "Invalid usage: missing command\n\nUsage:\n!run <command>\n\nExample:\n!run echo hello"
			return msg
		}
		BLACKLIST_CMDS := []string{"rm -rf", "mkfs", "dd", "shutdown","test"}
		for _, cmd := range BLACKLIST_CMDS {
			if strings.Contains(argument, cmd) {
				return "Galat command dalta hai! Try Again"
			}
		}
		output, err := runLocalCommand(argument)
		if err != nil {
			fmt.Println("Error")
		}
		return output
	} else {
		return "Not Authorised"
	}
}

