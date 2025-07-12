package bot

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "github.com/umxrrs/StorySwarm/internal/db"
    "github.com/umxrrs/StorySwarm/internal/web"
    "github.com/bwmarrin/discordgo"
)

type Config struct {
    Token     string `json:"token"`
    WebPort   string `json:"web_port"`
}

func GetConfig() Config {
    data, err := ioutil.ReadFile("config/config.json")
    if err != nil {
        log.Fatal("Error reading config:", err)
    }
    var config Config
    json.Unmarshal(data, &config)
    return config
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
        return
    }
    HandleCommands(s, m)
}

func OnReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
    if r.UserID == s.State.User.ID {
        return
    }
    HandleVote(s, r)
}

func StartWebServer() {
    web.Start(GetConfig().WebPort)
}

func StartStory(channelID, theme, tone string) error {
    return db.SaveStory(channelID, theme, tone)
}
