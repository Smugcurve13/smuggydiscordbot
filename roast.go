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
	userSplice := []string{}

	for _, msg := range userMessages {
		if msg.Author != nil && len(userSplice) < noOfMessages {
			userSplice = append(userSplice, msg.Content) 
		}
	}
	u := UserMessage{
		UserID: message.Author.ID,
		Username: message.Author.Username,
		Message: userSplice,
	}
	return u
}