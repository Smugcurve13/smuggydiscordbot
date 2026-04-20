package main

import (
	"fmt"
	"time"
	"os"

	"github.com/bwmarrin/discordgo"
)

func startAlertMonitor(session *discordgo.Session) {
	channel_id := os.Getenv("ALERT_CHANNEL_ID")
	for {
		stats,err := getStats()
		if err != nil {
			fmt.Printf("Error while fetching Stats : %v", err)
		}
		fmt.Println(stats,channel_id)
		time.Sleep(time.Minute * 5)
	}
}
