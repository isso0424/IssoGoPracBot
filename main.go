package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	stopBot = make(chan bool)
)

func main() {
	discord, err := discordgo.New()
	token := loadTokenFromEnv()
	if err != nil {
		fmt.Println(err)
	}
	discord.Token = "Bot " + token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(onReady)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	<-stopBot
}

func onReady(session *discordgo.Session, ready *discordgo.Ready) {
	sendMessage(session, "691261912008360000", "Hello world")
}

func sendMessage(session *discordgo.Session, channelID string, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		panic(err)
	}
}

func loadTokenFromEnv() string {
	fp, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var token string
	for scanner.Scan() {
		token = scanner.Text()
	}
	return token
}
