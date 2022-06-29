package commands

import (
	"log"
	"strings"

	"github.com/EdwardAndrew/MovieBuff/config"
	"github.com/EdwardAndrew/MovieBuff/pkg/api"

	"github.com/bwmarrin/discordgo"
)

func removeBotPrefix(content string, botId string) string {
	command := strings.TrimPrefix(content, config.Get().Prefix+" ")
	command = strings.TrimPrefix(command, "<@"+botId+">")
	command = strings.TrimSpace(command)
	return command
}

func Search(s *discordgo.Session, m *discordgo.Message) {
	term := removeBotPrefix(m.Content, s.State.User.ID)
	_, err := api.Search(term)
	if err != nil {
		log.Println("Something went wrong when calling the api...")
	}
}
