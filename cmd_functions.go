package main

import (
	"fmt"
	"strings"
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

func helpFunc(message *discordgo.MessageCreate, arg string) string {
	return "Available commands: !ping, !help, !stats and !run <command> <argument>"
}

func pingFunc(message *discordgo.MessageCreate, arg string) string {
	return `Pong ! Type "!help" to get all available commands`
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

func roastFunc(message *discordgo.MessageCreate, lastRoast string) string {
	roasts := []string{"I’d agree with you, but then we’d both be wrong.",
						"You’re the reason the gene pool needs a lifeguard.",
						"I’ve been called worse by people who are much better.",
						"You have the perfect face for radio and a great voice for silent films.",
						"I forgot the world revolves around you. My apologies—how silly of me to think other people existed."}

	idx := rand.IntN(len(roasts))
	selection := roasts[idx]
	if selection == lastRoast {
		idx = (idx - 1) % len(roasts)
		selection = roasts[idx]
	} 
	return selection
}

func quizFunc(message *discordgo.MessageCreate, arg string) string {
	// userId := message.Author.ID
	db := QuizDB{
		question: "What is smuggy's favourite number",
		options:[]string{"420","67","69","786"},
		answer_id: 0,
	}
	selection := 0
	if selection == db.answer_id {
		return "incoming"
	} else {
		return "You are not a quiz"
	}
}
