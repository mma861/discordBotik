package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ifInStr(str string, arr []string) string {
	for _, value := range arr {
		if strings.Contains(str, value) {
			return value
		}
	}
	return "FNORD"
}

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Не работае:", err)
		return
	}

	dg.AddHandler(messageReact)
	dg.AddHandler(messageReply)
	err = dg.Open()
	if err != nil {
		fmt.Println("Нет соединения:", err)
		return
	}
	defer dg.Close()

	fmt.Println("online")
	select {}
}

func messageReact(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	pregnant, err := readFile("pregnant.txt")
	if err != nil {
		fmt.Println("Нет файла:", err)
	}
	content := strings.ToLower(m.Content)
	// fmt.Println(strings.Contains(content, ifInStr(content, pregnant)))
	// fmt.Println(ifInStr(content, pregnant))
	if strings.Contains(content, ifInStr(content, pregnant)) {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "🫃")
		if err != nil {
			fmt.Println("Нет реакции:", err, "🫃")
		}
	}
}
func messageReply(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(m.Content, "🫃") {
		s.ChannelMessageSendReply(m.ChannelID, "🫃", m.Reference())
	}
}
