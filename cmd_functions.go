package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/genai"
)

func helpFunc(session *discordgo.Session, message *discordgo.MessageCreate, arg string) string {
	return "Available commands: !roast, !ping, !help, !stats and !run <command> <argument>"
}

func pingFunc(session *discordgo.Session, message *discordgo.MessageCreate, arg string) string {
	return `Pong ! Type "!help" to get all available commands`
}

func statsFunc(session *discordgo.Session, message *discordgo.MessageCreate, arg string) string {
	stats, err := getStats()
	if err != nil {
		return "Unable to fetch system stats"
	}
	return stats
}

func runFunc(session *discordgo.Session, message *discordgo.MessageCreate, argument string) string {
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

func roastFuncv2(session *discordgo.Session, message *discordgo.MessageCreate, argument string) string {
	targetUserID, err := getRoastTargetUser(message, argument)
	if err != nil {
		return "Custom Arguments not allowed as of now"
	}
	msgStruct := fetchMessagesofUserID(session, message, targetUserID, 6)
	msgs := []string{}
	msgs = msgStruct.Message
	msgs2 := strings.Join(msgs, ", ")
	result := aiRoast(msgs2)
	// result := testingaiRoast(msgs2)
	return result 
}

func quizFunc(session *discordgo.Session, message *discordgo.MessageCreate, arg string) string {
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

func testaiFunc(session *discordgo.Session, message *discordgo.MessageCreate, arg string) string {
	fmt.Printf("this is a test msg")
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
		apiKey := os.Getenv("GEMINI_API_KEY")
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		client := geminiClient(ctx, apiKey)

		parts := []*genai.Part{
			{Text: arg},
		}
		contents := []*genai.Content{{Parts: parts}}

		response, err := client.Models.GenerateContent(ctx, geminiModel, contents, nil)
		if err != nil {
			fmt.Printf("GenerateContent Error : %s" , err)
			return "Please try again Later , Model is Overloaded right now"
		}
		clean_response := cleanGeminiResponse(response)
		return clean_response
	} else {
		return "Not Authorized"
	}
}