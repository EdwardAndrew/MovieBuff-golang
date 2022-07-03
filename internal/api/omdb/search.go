package omdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/EdwardAndrew/MovieBuff/internal/api"
	"github.com/EdwardAndrew/MovieBuff/internal/config"
	"github.com/bwmarrin/discordgo"
)

func (o OMDB) Search(term string) (api.CachedSearchAPIResponse, error) {
	result := api.CachedSearchAPIResponse{
		SearchKey: "",
		Data:      "",
		HasData:   false,
	}
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

	if resp.StatusCode != http.StatusOK {
		return result, errors.New("Did not receive StatusOK from OMDB")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	json := new(OMDBSearchResult)
	json.FromString(string(data))

	if json.Response != "False" {
		return api.CachedSearchAPIResponse{
			HasData:   true,
			Data:      string(data),
			SearchKey: json.Title,
		}, nil
	}

	return result, nil
}

func (o OMDB) GetMessageEmbedFromData(data string) (*discordgo.MessageEmbed, error) {
	osr := new(OMDBSearchResult)
	err := osr.FromString(data)
	return omdbResponseToMessageEmbed(osr), err
}

func (o *OMDBSearchResult) FromString(s string) error {
	err := json.Unmarshal([]byte(s), o)
	return err
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
