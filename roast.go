package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/genai"
)

func fetchMessagesofUserID(session *discordgo.Session, message *discordgo.MessageCreate, targetUserID string, noOfMessages int) UserMessage {
	var userMessages []*discordgo.Message
	messages, err := session.ChannelMessages(message.ChannelID, 100, "", "", "")
	if err != nil {
		return UserMessage{}
	}

	for _, m := range messages {
		if m.Author.ID == targetUserID {
			userMessages = append(userMessages, m)
		}
	}
	messageSplice := []string{}

	for _, msg := range userMessages {
		if msg.Author != nil && len(messageSplice) < noOfMessages {
			messageSplice = append(messageSplice, msg.Content) 
		}
	}
	u := UserMessage{
		UserID: message.Author.ID,
		Username: message.Author.Username,
		Message: messageSplice,
	}
	return u
}

func aiRoast(msgs string) string {
	apiKey := os.Getenv("GEMINI_API_KEY")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client := geminiClient(ctx, apiKey)

	prompt := fmt.Sprintf(
		`Roast the user based on their recent Discord messages.

		Rules:
		- Be funny, not abusive
		- Keep it short (2–3 lines)
		- Focus on patterns in messages

		Messages:\n
		%s`, msgs)

	parts := []*genai.Part{
		{Text: prompt},
	}
	contents := []*genai.Content{{Parts: parts}}

	response, err := client.Models.GenerateContent(ctx, "gemini-flash-lite-latest", contents, nil)
	if err != nil {
		fmt.Printf("GenerateContent Error : %s" , err)
		return "Please try again Later , Model is Overloaded right now"
	}
	clean_response := cleanGeminiResponse(response)
	return clean_response
}