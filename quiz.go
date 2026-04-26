package main

// import (
// 	"github.com/bwmarrin/discordgo"
// )

// func dropdown(session *discordgo.Session, message *discordgo.Message) {

// 	menu := &discordgo.SelectMenu{
// 		CustomID:    "option_sele",
// 		Placeholder: "Choose an option",
// 		Options: []discordgo.SelectMenuOption{
// 			{
// 				Label:       "Label1",
// 				Value:       "Value1",
// 				Description: "Desc1",
// 			},
// 			{
// 				Label:       "Label2",
// 				Value:       "Value2",
// 				Description: "Desc2",
// 			},
// 		},
// 	}
// 	row := &discordgo.ActionsRow{
// 		Components: []discordgo.MessageComponent{menu},
// 	}

// 	session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
// 		Content:    "Select bitch",
// 		Components: []discordgo.MessageComponent{row},
// 	})

// }