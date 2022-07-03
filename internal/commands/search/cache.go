package search

import (
	"errors"
	"log"
	"strconv"
	"strings"

	a "github.com/EdwardAndrew/MovieBuff/internal/api"
	"github.com/EdwardAndrew/MovieBuff/pkg/cache"
	"github.com/bwmarrin/discordgo"
)

func fetchDataFromCache(api a.CachedSearchAPI, search string, incrmentCount bool) (bool, string, error) {
	hasSearchKey, searchKey, err := cache.Get(strings.Join([]string{api.GetSearchCachePrefix(), search}, ""))
	if err != nil {
		return false, "", err
	}

	if hasSearchKey {
		hasCachedData, cachedData, err := cache.Get(strings.Join([]string{api.GetDataCachePrefix(), searchKey}, ""))
		if err != nil {
			return false, "", err
		}

		if hasCachedData {
			if incrmentCount {
				incrementSearchCount(searchKey, api)
			}
			return true, cachedData, err
		}
	}

	return false, "", nil
}

func cacheResponse(search string, api a.CachedSearchAPI, result a.CachedSearchAPIResponse) {
	keys := []string{api.GetSearchCachePrefix() + search, api.GetDataCachePrefix() + result.SearchKey}
	values := []string{result.SearchKey, result.Data}

	err := cache.SetMultiple(keys, values)
	if err != nil {
		log.Print(err)
	} else {
		log.Println("Stored in cache.")
	}
}

func incrementSearchCount(cacheKey string, api a.CachedSearchAPI) error {
	return cache.Increment(api.GetCountCachePrefix() + cacheKey)
}

func setAskedCountFooter(embed *discordgo.MessageEmbed, cacheKey string, api a.CachedSearchAPI) *discordgo.MessageEmbed {
	askedBeforeCount, err := getSearchCount(cacheKey, api)
	if err != nil {
		return embed
	}

	var footerText string

	if askedBeforeCount <= 1 {
		footerText = "You're the first person to ask about this."
	} else {
		footerText = "Asked " + strconv.Itoa(askedBeforeCount) + " times before."
	}
	embed.Footer = &discordgo.MessageEmbedFooter{Text: footerText}
	return embed
}

func getSearchCount(cacheKey string, api a.CachedSearchAPI) (int, error) {
	found, stringValue, err := cache.Get(api.GetCountCachePrefix() + cacheKey)
	if err != nil {
		return -1, err
	}
	if !found {
		return -1, errors.New("searchCount: failed to find value for key")
	}

	numberValue, err := strconv.Atoi(stringValue)
	if err != nil {
		return -1, err
	}

	return numberValue, nil
}
