package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/EdwardAndrew/MovieBuff/pkg/handlers"
	"github.com/bwmarrin/discordgo"
)

func start() {
	discord, err := discordgo.New("Bot " + "NzcxMTI4NDE0Nzk5MTM0NzIw.GNHtYo.11J11Psx0InfeJIzSuiP7xhQueFYLrGhESyT8Y")

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

	log.Println("Bot started.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func main() {
	start()
}
