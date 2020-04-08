package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	stopBot = make(chan bool)
)

func main() {
	discord, err := discordgo.New()
	discord.Token = Token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	<-stopBot
}

func onMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	c, err := session.State.Channel(message.ChannelID)
}
