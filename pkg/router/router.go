package router

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type String string

func (s *String) tolower() *String {
	*s = String(strings.ToLower(string(*s)))
	return s
}

func (s *String) toupper() *String {
	*s = String(strings.ToUpper(string(*s)))
	return s
}

var helpCommands = map[string]bool{
	"help":     true,
	"h":        true,
	"commands": true,
	"i":        true,
	"info":     true,
}

func Route(message *discordgo.Message) {
	//TODO: route messages

	switch "test" {
	case "h":
	case "help":
	case "info":
	case "commands":
	default:
	}

}
