package router

import (
	"github.com/EdwardAndrew/MovieBuff/internal/commands/search"
	"github.com/EdwardAndrew/MovieBuff/internal/config"
	u "github.com/EdwardAndrew/MovieBuff/internal/utils"
	"github.com/bwmarrin/discordgo"
)

func Route(s *discordgo.Session, m *discordgo.Message) {
	cmd := u.RemoveBotPrefix(m.Content, s.State.User.ID)
	switch {
	case cmd == "h", cmd == "help", cmd == "info", cmd == "commands":
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: "I'd love to help you...",
			Embeds:  []*discordgo.MessageEmbed{u.GetHelpEmbed(s)},
		})
	case cmd == "prefix":
		s.ChannelMessageSend(m.ChannelID, u.FormatResponse(config.Get().Prefix+" is my prefix."))
	default:
		search.Search(s, m)
	}
}
