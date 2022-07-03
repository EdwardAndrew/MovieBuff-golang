package omdb

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/EdwardAndrew/MovieBuff/internal/config"
	"github.com/bwmarrin/discordgo"
)

func (o OMDB) Search(term string) (*discordgo.MessageEmbed, error) {
	result := new(discordgo.MessageEmbed)

	req, err := http.NewRequest("GET", o.baseURL, nil)
	if err != nil {
		return result, err
	}

	q := req.URL.Query()
	q.Add("t", term)
	q.Add("plot", "full")
	q.Add("y", "")
	q.Add("apikey", config.Get().OMDB.API_Key)

	req.URL.RawQuery = q.Encode()

	resp, err := o.client.Get(req.URL.String())
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	log.Println("Received response from OMDB.")

	if resp.StatusCode != http.StatusOK {
		return result, errors.New("Did not receive StatusOK from OMDB")
	}

	omdbResponse := new(OMDBSearchResult)
	err = json.NewDecoder(resp.Body).Decode(&omdbResponse)
	if err != nil {
		return result, err
	}

	return omdbResponseToMessageEmbed(omdbResponse), nil
}

func omdbResponseToMessageEmbed(o *OMDBSearchResult) *discordgo.MessageEmbed {
	embed := new(discordgo.MessageEmbed)

	embed.Title = o.Title
	embed.Description = o.Plot
	embed.Color = 16224842 //#f7924a
	embed.Footer = &discordgo.MessageEmbedFooter{Text: "asked before amount"}

	if o.Poster != "N/A" {
		embed.Image = &discordgo.MessageEmbedImage{URL: o.Poster}
	}
	embed.URL = "https://www.imdb.com/title/" + o.IMDBId
	embed.Fields = append(embed.Fields,
		&discordgo.MessageEmbedField{
			Name:  "Actors",
			Value: o.Actors,
		})

	return embed
}