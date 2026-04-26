package main

import (

	"github.com/bwmarrin/discordgo"
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