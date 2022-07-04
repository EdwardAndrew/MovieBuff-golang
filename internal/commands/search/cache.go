package search

import (
	"errors"
	"log"
	"strconv"
	"strings"

	a "github.com/EdwardAndrew/MovieBuff/internal/api"
	"github.com/EdwardAndrew/MovieBuff/pkg/cache"
)

type cacheFetchResponse struct {
	Found    bool
	Data     string
	CacheKey string
}

func cacheFetch(api a.CachedSearchAPI, search string) (cacheFetchResponse, error) {
	hasSearchKey, searchKey, err := cache.Get(strings.Join([]string{api.GetSearchCachePrefix(), search}, ""))
	if err != nil {
		return cacheFetchResponse{}, err
	}

	if hasSearchKey {
		hasCachedData, cachedData, err := cache.Get(strings.Join([]string{api.GetDataCachePrefix(), searchKey}, ""))
		if err != nil {
			return cacheFetchResponse{}, err
		}

		if hasCachedData {
			return cacheFetchResponse{Found: true, Data: cachedData, CacheKey: searchKey}, err
		}
	}

	return cacheFetchResponse{}, nil
}

func cacheQuery(search string, api a.CachedSearchAPI, result a.CachedSearchAPIResponse) {
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
