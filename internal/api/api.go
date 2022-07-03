package api

import "github.com/bwmarrin/discordgo"

type CachedSearchAPI interface {
	Search(term string) (CachedSearchAPIResponse, error)
	GetSearchCachePrefix() string
	GetDataCachePrefix() string
	GetCountCachePrefix() string
	GetName() string
	GetMessageEmbedFromData(data string) (*discordgo.MessageEmbed, error)
}

type CachedSearchAPIResponse struct {
	SearchKey, Data string
	HasData         bool
}
