package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"os/signal"
	"syscall"
	"strings"

	"github.com/bwmarrin/discordgo"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	token := os.Getenv("DISCORD_TOKEN")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	discord.Identify.Intents = discordgo.IntentsGuildMessages
	discord.AddHandler(MessageHandler)
	discorderr := discord.Open()
	if discorderr != nil {
		fmt.Printf("Error in opening Discord Session : %v", discorderr)
		os.Exit(1)
	}

	fmt.Println("Smuggy Bot is Running")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	discord.Close()
	fmt.Println("Smuggy Bot is shutting down")

}

func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	fmt.Println("Message Received")
	if message.Author.ID == session.State.User.ID {
		return
	}
	msg := message.Content
	if strings.HasPrefix(msg, "!") {
		switch msg {
		case "!ping" :
			session.ChannelMessageSend(message.ChannelID, "This is ping function")
		case "!stats" :
			session.ChannelMessageSend(message.ChannelID, "This is stats function")
		case "!help" :
			session.ChannelMessageSend(message.ChannelID, "This is help function")
		}
	}
}