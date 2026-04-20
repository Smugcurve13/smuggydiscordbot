package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"os/signal"
	"syscall"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/crypto/ssh"
)



func main() {
	fmt.Println(time.Now())
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
			stats, err := getStats()
			if err != nil {
				session.ChannelMessageSend(message.ChannelID, "Error Getting Stats")
				return
			}
			session.ChannelMessageSend(message.ChannelID, stats)
		case "!help" :
			session.ChannelMessageSend(message.ChannelID, "This is help function")
		}
	}
}

func getStats() (string, error) {
	host := os.Getenv("SSH_HOST")
	user := os.Getenv("SSH_USER")
	pass := os.Getenv("SSH_PASS")
	config := ssh.ClientConfig{
		User:	user,
		Auth:	[]ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host+":22", &config)
	if err != nil {
		fmt.Printf("Error Ocurred: %v", err)
	}
	
	defer client.Close()
	fmt.Println("SSH CONNECTED !!")
	return "" , nil
}