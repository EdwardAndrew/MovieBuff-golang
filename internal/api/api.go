package api

import (
	"github.com/bwmarrin/discordgo"
)

type CachedSearchAPI interface {
	Search(term string) (*discordgo.MessageEmbed, error)
	GetSearchCachePrefix() string
	GetDataCachePrefix() string
	GetCountCachePrefix() string
	GetName() string
}
