package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	token := os.Getenv("DISCORD_TOKEN")
	discord,err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	fmt.Println(discord)
}