package handlers

import (
	"log"

	"github.com/EdwardAndrew/MovieBuff/internal/config"
	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Println(config.Get().BotName + " " + config.Get().Version + " ready.")
}
