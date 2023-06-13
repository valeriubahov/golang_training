package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func printList(s *discordgo.Session, channelID string) {
	for i, v := range toDoList {
		value := fmt.Sprint(i+1) + " - " + string(v.name)
		s.ChannelMessageSend(channelID, value)
	}
}
