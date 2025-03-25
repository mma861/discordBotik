package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Не работае:", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Нет соединения:", err)
		return
	}
	defer dg.Close()

	fmt.Println("online")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.ToLower(m.Content)
	if strings.Contains(content, "жоско") || strings.Contains(content, "жоска") {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "🫃")
		if err != nil {
			fmt.Println("Нет реакции:", err)
		}
	}
}
