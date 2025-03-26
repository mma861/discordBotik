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
	pregFile := "pregnant.txt"
	swearFile := "swears.txt"
	govnoFile := "govno.txt"

	pregnantReact, err := readFile(pregFile)
	if err != nil {
		fmt.Println("Нет файла:", err)
	}
	deleteMsg, err := readFile(swearFile)
	if err != nil {
		fmt.Println("Нет файла:", err)
	}
	govnoMsg, err := readFile(govnoFile)
	if err != nil {
		fmt.Println("Нет файла:", err)
	}

	content := strings.ToLower(m.Content)
	// fmt.Println(strings.Contains(content, ifInStr(content, pregnant)))
	// fmt.Println(ifInStr(content, pregnant))
	if strings.Contains(content, ifInStr(content, pregnantReact)) || strings.Contains(content, ifInStr(content, govnoMsg)) {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "🫃")
		if err != nil {
			fmt.Println("Нет реакции:", err, "🫃")
		}
	}

	if strings.Contains(content, ifInStr(content, deleteMsg)) {
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			fmt.Println("Нет удаления:", err, m.ChannelID, m.ID)
		}
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(content, "🫃") {
		s.ChannelMessageSendReply(m.ChannelID, "🫃", m.Reference())
	}

	if strings.Contains(content, ifInStr(content, govnoMsg)) {
		s.ChannelMessageSendReply(m.ChannelID, ifInStr(content, govnoMsg)+" говно", m.Reference())
	}
}
