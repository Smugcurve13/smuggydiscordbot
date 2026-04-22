package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func helpFunc(arg string) string {
	return "This is help function"
}

func pingFunc(arg string) string {
	return "This is ping function"
}

func statsFunc(arg string) string {
	stats, err := getStats()
	if err != nil {
		return "Error Getting Stats"
	}
	return stats
}

func runFunc(session *discordgo.Session, message *discordgo.MessageCreate, argument string) {
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

