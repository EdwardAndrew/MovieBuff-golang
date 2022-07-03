package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type omdb struct {
	API_Key  string
	Base_URL string
}

type config struct {
	Prefix            string
	OMDB              omdb
	DiscordToken      string
	BotName           string
	Version           string
	DeveloperId       string
	CacheGlobalPrefix string
}

type environmentVariable struct {
	value string
	key   string
}

var hasBeenLoaded bool
var c config

func loadConfig() config {
	omdbAPIKey, present := os.LookupEnv("OMDB_API_KEY")
	if !present {
		log.Fatal("OMDB API Key is not set.")
	}

	discordToken, present := os.LookupEnv("DISCORD_TOKEN")
	if !present {
		log.Fatal("Discord Token is not set.")
	}

	developerId, present := os.LookupEnv("DEVELOPER_ID")
	if !present {
		log.Fatal("Developer Id is not set.")
	}

	result := config{
		Prefix: "!mb",
		OMDB: omdb{
			API_Key:  omdbAPIKey,
			Base_URL: "https://www.omdbapi.com",
		},
		DiscordToken:      discordToken,
		BotName:           "MovieBuff",
		DeveloperId:       developerId,
		CacheGlobalPrefix: "moviebuff/",
		Version:           "2.0.0",
	}

	return result
}

func Get() config {
	if !hasBeenLoaded {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file of config")
		}
		c = loadConfig()

		hasBeenLoaded = true
	}
	return c
}
