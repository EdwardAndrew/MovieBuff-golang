package handlers

import (
	"strings"

	"github.com/EdwardAndrew/MovieBuff/config"
	"github.com/EdwardAndrew/MovieBuff/pkg/router"
	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore from itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore if string doesn't start with bot prefix
	if !strings.HasPrefix(m.Content, config.Config.Prefix) {
		return
	}

	router.Route(m.Message)
}
