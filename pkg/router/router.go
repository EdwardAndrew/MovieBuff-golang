package router

import (
	"github.com/EdwardAndrew/MovieBuff/pkg/commands"
	"github.com/EdwardAndrew/MovieBuff/pkg/config"
	u "github.com/EdwardAndrew/MovieBuff/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

var helpEmbed = discordgo.MessageEmbed{
	Title:       config.Get().BotName + " Help",
	Description: "Here's some example of how to use " + config.Get().BotName + "!",
	Author:      &discordgo.MessageEmbedAuthor{Name: config.Get().BotName},
	Fields: []*discordgo.MessageEmbedField{
		{Name: "Anime Search", Value: config.Get().Prefix + " -anime Toradora!"},
		{Name: "Developer", Value: config.Get().DeveloperId},
		{Name: "Version", Value: config.Get().Version},
	},
}

func Route(s *discordgo.Session, m *discordgo.Message) {
	command := u.RemoveBotPrefix(m.Content, s.State.User.ID)

	switch command {
	case "h":
		fallthrough
	case "help":
		fallthrough
	case "info":
		fallthrough
	case "commands":
		s.ChannelMessageSendEmbed(m.ChannelID, &helpEmbed)
	case "prefix":
		s.ChannelMessageSend(m.ChannelID, u.FormatResponse(config.Get().Prefix+" is my prefix."))
	default:
		commands.Search(s, m)
	}

}
