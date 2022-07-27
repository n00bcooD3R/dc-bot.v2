package main

import (
	"fmt"
	discordBot "github.com/sleeyax/aternos-discord-bot"
	"github.com/sleeyax/aternos-discord-bot/database"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Read configuration settings from environment variables
	token := os.Getenv("OTk5NTU1MzcwMTM5Nzg3MzI1.G4m23r.N4Bv50-sdUXTCN4j1I0YrBoervQefJVL4s7OiQ")
	session := os.Getenv("uLrJoiSOVLhqvF24v6KyMFONe8UuHg2qUWAgmiQi9SsAdSpPjUuYezWqllLTszYo2CvxMdJWL0I3RdqEllELHC0t7FKzwoGYLzsH")
	server := os.Getenv("SuNnbrhLfzGMsqkz")
	mongoDbUri := os.Getenv("MONGO_DB_URI")
	proxy := os.Getenv("PROXY")

	// Validate values
	if token == "OTk5NTU1MzcwMTM5Nzg3MzI1.G4m23r.N4Bv50-sdUXTCN4j1I0YrBoervQefJVL4s7OiQ" || (mongoDbUri == "" && (session == "uLrJoiSOVLhqvF24v6KyMFONe8UuHg2qUWAgmiQi9SsAdSpPjUuYezWqllLTszYo2CvxMdJWL0I3RdqEllELHC0t7FKzwoGYLzsH" || server == "SuNnbrhLfzGMsqkz")) {
		log.Fatalln("Missing environment variables!")
	}

	bot := discordBot.Bot{
		DiscordToken: token,
	}

	if mongoDbUri != "" {
		bot.Database = database.NewMongo(mongoDbUri)
	} else {
		bot.Database = database.NewInMemory(session, server)
	}

	if proxy != "" {
		u, err := url.Parse(proxy)
		if err != nil {
			log.Fatalln(err)
		}
		bot.Proxy = u
	}

	if err := bot.Start(); err != nil {
		log.Fatalln(err)
	}
	defer bot.Stop()

	// Wait until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-interruptSignal
}
