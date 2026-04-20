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
	if message.Author.ID == session.State.User.ID {
		return
	}
	fmt.Println("Message Received")
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

	ram_cmd := "free -m"
	cpu_cmd := `grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$3+$4+$5)} END {print usage"%"}'`
	disk_cmd := "df -h /"
	temp_cmd := `sensors | grep "Package id 0" | awk '{print $4}'`
	
	ram , err1 := runCommand(client, ram_cmd)
	if err1 != nil {
		fmt.Println("Error while excecuting RAM stats command")
	}
	cpu , err2 := runCommand(client, cpu_cmd)
	if err2 != nil {
		fmt.Println("Error while excecuting CPU stats command")
	}
	disk , err3 := runCommand(client, disk_cmd)
	if err3 != nil {
		fmt.Println("Error while excecuting DISK stats command")
	}
	temp , err4 := runCommand(client, temp_cmd)
	if err4 != nil {
		fmt.Println("Error while excecuting TEMPERATURE stats command")
	}
	result := fmt.Sprintf("```\nSmuggyServer Stats\n\nRAM:\n%v\nCPU:\n%v\nDISK:\n%v\nTEMPERATURE:\n%v\n```", ram, cpu, disk, temp)
	fmt.Println(result)
	return result , nil
}

func runCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Error in Creating New Session : %v", err)
	}
	defer session.Close()
	output, err3 := session.Output(command)
	if err3 != nil {
		fmt.Printf("Error in Running Command : %v", err3)
	}
	str_output := string(output)
	return str_output, nil
}