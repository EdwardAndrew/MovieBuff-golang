package search

import (
	"log"
	"strconv"

	a "github.com/EdwardAndrew/MovieBuff/internal/api"
	"github.com/EdwardAndrew/MovieBuff/internal/api/omdb"
	u "github.com/EdwardAndrew/MovieBuff/internal/utils"

	"github.com/bwmarrin/discordgo"
)

func Search(s *discordgo.Session, m *discordgo.Message) {
	searchTerm := u.RemoveBotPrefix(m.Content, s.State.User.ID)

	resp, err := query(searchTerm)
	if err != nil {
		log.Println("Something went wrong when calling the OMDB api.")
		return
	}
	if !resp.Found {
		s.MessageReactionAdd(m.ChannelID, m.ID, "😭")
		s.ChannelMessageSend(m.ChannelID, u.FormatResponse("I couldn't find anything, sorry!"))
		return
	}

	reply := &discordgo.MessageSend{
		Embeds:  []*discordgo.MessageEmbed{resp.Embed},
		Content: u.FormatResponse("Here's what I can tell you about _" + resp.Embed.Title + "_"),
	}
	s.ChannelMessageSendComplex(m.ChannelID, reply)
}

func getRelevantAPI() a.CachedSearchAPI {
	return omdb.API
}

type queryResponse struct {
	Found bool
	Embed *discordgo.MessageEmbed
}

func apiFetch(api a.CachedSearchAPI, search string, storeResultInCache bool) (queryResponse, error) {
	result, err := api.Search(search)
	if err != nil {
		return queryResponse{}, err
	}

	if !result.HasData {
		return queryResponse{}, nil
	}

	embed, err := api.GetMessageEmbedFromData(result.Data)
	if err != nil {
		return queryResponse{}, err
	}

	if storeResultInCache {
		go cacheQuery(search, api, result)
	}
	return queryResponse{Found: true, Embed: embed}, nil
}

func query(search string) (queryResponse, error) {
	var embed *discordgo.MessageEmbed = nil
	api := getRelevantAPI()
	forceAPIFetch := false

	cacheFetchResponse, err := cacheFetch(api, search)
	if cacheFetchResponse.Found && err == nil {
		embed, err = api.GetMessageEmbedFromData(cacheFetchResponse.Data)
		log.Println("Retrived from cache.")
		if err != nil {
			log.Println("Error making embed from cached data.")
			forceAPIFetch = true
		}
	}

	if !cacheFetchResponse.Found || forceAPIFetch {
		apiResponse, err := apiFetch(api, search, true)
		log.Println("Fetched from server.")

		if apiResponse.Found && err == nil {
			embed = apiResponse.Embed
		} else {
			return queryResponse{}, err
		}
	}

	incrementSearchCount(embed.Title, api)
	setFooterAskedBeforeText(embed, api)

	return queryResponse{Found: true, Embed: embed}, nil
}

func setFooterAskedBeforeText(embed *discordgo.MessageEmbed, api a.CachedSearchAPI) error {
	searchCount, err := getSearchCount(embed.Title, api)
	if err != nil {
		log.Print(err)
		return err
	}

	var footerText string

	if searchCount == 1 {
		footerText = "You're the first person to ask about this."
	} else {
		footerText = "Asked " + strconv.Itoa(searchCount) + " times before."
	}
	embed.Footer.Text = footerText
	return nil
}
