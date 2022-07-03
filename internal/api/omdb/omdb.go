package omdb

import (
	"net/http"
	"time"

	"github.com/EdwardAndrew/MovieBuff/internal/config"
)

type OMDB struct {
	name, baseURL string
	client        http.Client
}

var API OMDB = OMDB{
	name:    "omdb",
	baseURL: config.Get().OMDB.Base_URL,
	client: http.Client{
		Timeout: time.Second * 10,
	},
}

func (o OMDB) GetName() string {
	return o.name
}

func (o OMDB) GetSearchCachePrefix() string {
	return config.Get().CacheGlobalPrefix + "movieSearch-omdb/"
}

func (o OMDB) GetDataCachePrefix() string {
	return config.Get().CacheGlobalPrefix + "movie-omdb/"
}

func (o OMDB) GetCountCachePrefix() string {
	return config.Get().CacheGlobalPrefix + "count/"
}
