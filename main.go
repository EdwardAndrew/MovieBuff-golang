package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/EdwardAndrew/MovieBuff/pkg/config"
	"github.com/EdwardAndrew/MovieBuff/pkg/handlers"
)

func main() {
	discord, err := discordgo.New("Bot " + config.Get().DiscordToken)

	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
		return
	}
	handlers.BindHandlers(discord)
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		log.Fatal("Error connecting to Discord: ", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}
