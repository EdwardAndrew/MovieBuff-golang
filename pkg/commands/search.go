package commands

import (
	"log"

	"github.com/EdwardAndrew/MovieBuff/pkg/api/omdb"
	u "github.com/EdwardAndrew/MovieBuff/pkg/utils"

	"github.com/bwmarrin/discordgo"
)

func Search(s *discordgo.Session, m *discordgo.Message) {
	term := u.RemoveBotPrefix(m.Content, s.State.User.ID)
	search := getRelevantSearch()

	embed, err := search(term)
	if err != nil {
		s.MessageReactionAdd(m.ChannelID, m.ID, "991756430900199565")
		s.ChannelMessageSend(m.ChannelID, u.FormatResponse("I couldn't find anything, sorry!"))

		log.Println("Something went wrong when calling the api...")
		return
	}

	reply := &discordgo.MessageSend{
		Embeds:  []*discordgo.MessageEmbed{embed},
		Content: u.FormatResponse("Here's what I can tell you about _" + embed.Title + "_"),
	}
	s.ChannelMessageSendComplex(m.ChannelID, reply)
}

func getRelevantSearch() func(string) (*discordgo.MessageEmbed, error) {
	return omdb.Search
}
