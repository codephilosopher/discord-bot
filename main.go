package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	discord_token := os.Getenv("DISCORD_TOKEN")
	sess, err := discordgo.New("Bot " + discord_token)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			s.ChannelMessageSend(m.ChannelID, "pong")
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	if err = sess.Open(); err != nil {
		log.Fatal(err)
	}
	defer sess.Close()
	log.Println("The Bot is online!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
