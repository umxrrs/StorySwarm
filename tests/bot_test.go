package bot

import (
    "testing"
    "github.com/StorySwarm/internal/db"
    "github.com/bwmarrin/discordgo"
)

func TestHandleCommands(t *testing.T) {
    db.Init()
    s := &discordgo.Session{}
    m := &discordgo.MessageCreate{
        Message: &discordgo.Message{
            Content:   "!story start fantasy epic",
            ChannelID: "test-channel",
            Author:    &discordgo.User{ID: "user1"},
        },
    }
    HandleCommands(s, m)
    story, err := db.GetActiveStory("test-channel")
    if err != nil || story == nil || story.Theme != "fantasy" || story.Tone != "epic" {
        t.Error("Failed to start story")
    }
}
