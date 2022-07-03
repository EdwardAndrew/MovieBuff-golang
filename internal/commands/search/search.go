package search

import (
	"log"
	"strings"

	a "github.com/EdwardAndrew/MovieBuff/internal/api"
	"github.com/EdwardAndrew/MovieBuff/internal/api/omdb"
	u "github.com/EdwardAndrew/MovieBuff/internal/utils"
	"github.com/EdwardAndrew/MovieBuff/pkg/cache"

	"github.com/bwmarrin/discordgo"
)

func Search(s *discordgo.Session, m *discordgo.Message) {
	searchTerm := u.RemoveBotPrefix(m.Content, s.State.User.ID)

	embed, err := fetchData(searchTerm)
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

func getRelevantAPI() a.CachedSearchAPI {
	return omdb.API
}

func fetchData(term string) (*discordgo.MessageEmbed, error) {
	api := getRelevantAPI()

	cacheKey, err := cache.Get(strings.Join([]string{api.GetSearchCachePrefix(), term}, ""))
	if err != nil {
		log.Println("Something went wrong")
	}

	//TODO: Try to fetch data from redis cache

	//TODO: If data exists we need to parse it back into a discord embed and return it

	//TODO: Use a longish (3 months maybe?) TTL for storing cached data.

	log.Print(cacheKey)

	return nil, nil
}
