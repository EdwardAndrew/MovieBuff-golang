package search

import (
	"log"

	a "github.com/EdwardAndrew/MovieBuff/internal/api"
	"github.com/EdwardAndrew/MovieBuff/internal/api/omdb"
	u "github.com/EdwardAndrew/MovieBuff/internal/utils"

	"github.com/bwmarrin/discordgo"
)

func Search(s *discordgo.Session, m *discordgo.Message) {
	searchTerm := u.RemoveBotPrefix(m.Content, s.State.User.ID)

	success, embed, err := fetchEmbed(searchTerm)

	if err != nil {
		log.Println("Something went wrong when calling the OMDB api.")
		return
	}
	if !success {
		s.MessageReactionAdd(m.ChannelID, m.ID, "😭")
		s.ChannelMessageSend(m.ChannelID, u.FormatResponse("I couldn't find anything, sorry!"))
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

func fetchEmbed(search string) (bool, *discordgo.MessageEmbed, error) {
	api := getRelevantAPI()

	// Attempt to use cached data
	didFindCachedData, data, err := fetchDataFromCache(api, search, true)
	if didFindCachedData {
		if embed, err := api.GetMessageEmbedFromData(data); err == nil {
			log.Println("Retrived from cache.")
			return true, embed, nil
		}
	}

	// Fetch new data from API
	result, err := api.Search(search)
	if err != nil {
		return false, nil, err
	}

	log.Print(result)
	if !result.HasData {
		return false, nil, nil
	}
	incrementSearchCount(result.SearchKey, api)

	embed, err := api.GetMessageEmbedFromData(result.Data)
	if err != nil {
		return false, nil, err
	}
	embed = setAskedCountFooter(embed, result.SearchKey, api)

	go cacheResponse(search, api, result)

	log.Println("Fetched from server.")
	return true, embed, nil
}
