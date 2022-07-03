package utils

import (
	"strings"

	"github.com/EdwardAndrew/MovieBuff/internal/config"
	"github.com/bwmarrin/discordgo"
)

const BlankChar = "\u200B"

func FormatResponse(message string) string {
	return BlankChar + message
}

func RemoveBotPrefix(content string, botId string) string {
	command := strings.TrimPrefix(content, config.Get().Prefix+" ")
	command = strings.TrimPrefix(command, "<@"+botId+">")
	command = strings.TrimSpace(command)
	return command
}

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

func GetHelpEmbed() *discordgo.MessageEmbed {
	return &helpEmbed
}
