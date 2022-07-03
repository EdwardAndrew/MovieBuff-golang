package handlers

import (
	"strings"

	"github.com/EdwardAndrew/MovieBuff/internal/config"
	"github.com/EdwardAndrew/MovieBuff/internal/router"
	"github.com/bwmarrin/discordgo"
)

func containsAuthor(s []*discordgo.User, e *discordgo.User) bool {
	for _, user := range s {
		if user.ID == e.ID {
			return true
		}
	}
	return false
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore from itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore if string doesn't start with bot prefix or mention the bot
	if !strings.HasPrefix(m.Content, config.Get().Prefix) && !containsAuthor(m.Mentions, s.State.User) {
		return
	}

	router.Route(s, m.Message)
}
