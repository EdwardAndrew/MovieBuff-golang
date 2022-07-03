package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/EdwardAndrew/MovieBuff/internal/config"
	"github.com/EdwardAndrew/MovieBuff/internal/handlers"
	"github.com/EdwardAndrew/MovieBuff/pkg/cache"
	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v9"
)

func main() {
	cache.Connect(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // default
	})
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
