package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type Todo struct {
	name string
	done bool
}

var toDoList = make([]Todo, 0)

type configStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}

func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Println(m.Content)

	switch {
	case strings.Contains(m.Content, "add"):
		c := strings.ReplaceAll(m.Content, "add ", "")
		s.ChannelMessageSend(m.ChannelID, "Added")
		toAdd := Todo{name: c, done: false}
		toDoList = append(toDoList, toAdd)
		printList(s, m.ChannelID)

	case strings.Contains(m.Content, "remove"):
		c, err := strconv.Atoi(strings.ReplaceAll(m.Content, "remove ", ""))
		if err != nil {
			fmt.Println("Error")
		}

		toDoList = append(toDoList[:c-1], toDoList[c:]...)
		printList(s, m.ChannelID)
	case strings.Contains(m.Content, "show"):
		s.ChannelMessageSend(m.ChannelID, "You have the following items in the list:")
		printList(s, m.ChannelID)

	case strings.Contains(m.Content, "ping"):
		link := strings.ReplaceAll(m.Content, "ping ", "")
		pingWebsites(s, m.ChannelID, link)

	case strings.Contains(m.Content, "bot"):
		s.ChannelMessageSend(m.ChannelID, "How can I help you?")
	}
}

func main() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	<-make(chan struct{})
}
