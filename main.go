package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
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
	discord.AddHandler(MessageHandlerv3)
	discorderr := discord.Open()
	if discorderr != nil {
		fmt.Printf("Error in opening Discord Session : %v", discorderr)
		os.Exit(1)
	}

	fmt.Println("Smuggy Bot is Running")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go startAlertMonitor(discord)
	<-sigs
	discord.Close()
	fmt.Println("Smuggy Bot is shutting down")
}
