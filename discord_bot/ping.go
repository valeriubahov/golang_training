package main

import (
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

func pingWebsites(s *discordgo.Session, channelId string, link string) {
	c := make(chan string)

	go checkLink(link, c, s, channelId)

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c, s, channelId)
		}(l)
	}
}

// Check if the link return an error or not
// error != nil means that an error happened and the website is not running correctly
func checkLink(link string, c chan string, s *discordgo.Session, channelId string) {
	_, err := http.Get(link)

	if err != nil {
		s.ChannelMessageSend(channelId, "<"+link+"> might be down!")

		c <- link
		return
	}

	s.ChannelMessageSend(channelId, "<"+link+"> is up and running!")

	c <- link
}
