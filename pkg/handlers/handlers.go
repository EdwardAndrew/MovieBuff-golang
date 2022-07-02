package handlers

import "github.com/bwmarrin/discordgo"

func BindHandlers(s *discordgo.Session) {
	s.AddHandler(messageCreate)
	s.AddHandler(ready)
}
