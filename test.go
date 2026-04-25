package main // Changed from 'test' to 'main' so it runs

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/bwmarrin/discordgo"
// 	"github.com/joho/godotenv"
// )

// func dropdown(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	menu := discordgo.SelectMenu{
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

// 	// Important: Components must be inside an ActionsRow
// 	session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
// 		Content: "Select something:",
// 		Components: []discordgo.MessageComponent{
// 			discordgo.ActionsRow{
// 				Components: []discordgo.MessageComponent{menu},
// 			},
// 		},
// 	})
// }

// func main() {
// 	fmt.Println(time.Now())
// 	godotenv.Load()

// 	token := os.Getenv("DISCORD_TOKEN")
// 	discord, _ := discordgo.New("Bot " + token)

// 	// Set intents so the bot can see messages
// 	discord.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentMessageContent

// 	// 1. ADD THE HANDLER HERE
// 	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
// 		// Ignore messages from the bot itself
// 		if m.Author.ID == s.State.User.ID {
// 			return
// 		}

// 		// If someone types !test, show the dropdown
// 		if m.Content == "!test" {
// 			dropdown(s, m)
// 		}
// 	})
// 	// Add this alongside your MessageCreate handler
// 	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
//     // Check if the interaction is from a Message Component (like our dropdown)
//     if i.Type == discordgo.InteractionMessageComponent {
        
//         // Check if it's our specific menu using the CustomID we set
//         if i.MessageComponentData().CustomID == "option_sele" {
            
//             // Get the value the user selected (it's a slice of strings)
//             selectedValue := i.MessageComponentData().Values[0]

//             // Respond to the user so the "Interaction Failed" goes away
//             s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
//                 Type: discordgo.InteractionResponseChannelMessageWithSource,
//                 Data: &discordgo.InteractionResponseData{
//                     Content: "You picked: " + selectedValue,
//                     Flags:   discordgo.MessageFlagsEphemeral, // Only the user can see this
//                 },
//             })
//         }
//     }
// })

// 	err := discord.Open()
// 	if err != nil {
// 		fmt.Printf("Error opening connection: %v", err)
// 		return
// 	}

// 	fmt.Println("Bot is now running. Press CTRL-C to exit. Type !test in Discord!")
	
// 	// Keep the program running
// 	select {}
// }