package router

import (
	"strings"

	"github.com/EdwardAndrew/MovieBuff/config"
	"github.com/EdwardAndrew/MovieBuff/pkg/commands"
	"github.com/EdwardAndrew/MovieBuff/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

var helpEmbed = discordgo.MessageEmbed{
	Title:       config.Get().BotName + " Help",
	Description: "Here's some example of how to use " + config.Get().BotName + "!",
	Author:      &discordgo.MessageEmbedAuthor{Name: config.Get().BotName},
	Fields: []*discordgo.MessageEmbedField{
		{Name: "Anime Search", Value: "!mb -anime Toradora!"},
		{Name: "Developer", Value: config.Get().DeveloperId},
		{Name: "Version", Value: config.Get().Version},
	},
}

func removeBotPrefix(content string, botId string) string {
	command := strings.TrimPrefix(content, config.Get().Prefix+" ")
	command = strings.TrimPrefix(command, "<@"+botId+">")
	command = strings.TrimSpace(command)
	return command
}

func Route(s *discordgo.Session, m *discordgo.Message) {
	command := removeBotPrefix(m.Content, s.State.User.ID)

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
		s.ChannelMessageSend(m.ChannelID, utils.BlankChar+config.Get().Prefix+" is my prefix.")
	default:
		// s.ChannelMessageSend(m.ChannelID, "I don't understand that, sorry!")
		// s.MessageReactionAdd(m.ChannelID, m.ID, "991756430900199565")
		commands.Search(s, m)
	}

}
