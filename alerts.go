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
		stats, err := getRawStats()
		if err != nil {
			fmt.Printf("Error while fetching raw Stats : %v", err)
		}
		fmt.Printf("CPU: %.2f%% | RAM: %.2f%% | Disk: %.2f%% | Temp: %.2f°C\n", stats.CPUPercent, stats.RAMUsedPercent, stats.DiskUsedPercent, stats.TempCelsius)
		if stats.TempCelsius > 80 {
			session.ChannelMessageSend(channel_id, fmt.Sprintf("⚠️ HIGH TEMP ALERT: %.2f°C",stats.TempCelsius))
		}
		if stats.CPUPercent > 80 {
			session.ChannelMessageSend(channel_id, fmt.Sprintf("⚠️ HIGH CPU ALERT: %.2f%%",stats.CPUPercent))
		}
		if stats.RAMUsedPercent > 80 {
			session.ChannelMessageSend(channel_id, fmt.Sprintf("⚠️ HIGH RAM ALERT: %.2f%%",stats.RAMUsedPercent))
		}
		if stats.DiskUsedPercent > 80 {
			session.ChannelMessageSend(channel_id, fmt.Sprintf("⚠️ HIGH DISK ALERT: %.2f%%",stats.DiskUsedPercent))
		}
		time.Sleep(time.Minute * 5)
	}
}
