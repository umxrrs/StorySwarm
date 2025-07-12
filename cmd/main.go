package main

import (
    "log"
    "github.com/umxrrs/StorySwarm/internal/bot"
    "github.com/bwmarrin/discordgo"
)

func main() {
    dg, err := discordgo.New("Bot " + bot.GetConfig().Token)
    if err != nil {
        log.Fatal("Error creating Discord session:", err)
    }
    dg.AddHandler(bot.OnMessageCreate)
    dg.AddHandler(bot.OnReactionAdd)
    err = dg.Open()
    if err != nil {
        log.Fatal("Error connecting to Discord:", err)
    }
    log.Println("StorySwarm is online!")
    go bot.StartWebServer()
    select {}
}
