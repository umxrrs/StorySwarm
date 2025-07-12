package bot

import (
    "strings"
    "github.com/umxrrs/StorySwarm/internal/db"
    "github.com/umxrrs/StorySwarm/internal/sentiment"
    "github.com/bwmarrin/discordgo"
)

func HandleCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
    content := strings.TrimSpace(m.Content)
    if strings.HasPrefix(content, "!story start ") {
        parts := strings.SplitN(content, " ", 4)
        if len(parts) < 4 {
            s.ChannelMessageSend(m.ChannelID, "Usage: !story start <theme> <tone>")
            return
        }
        theme, tone := parts[2], parts[3]
        err := StartStory(m.ChannelID, theme, tone)
        if err != nil {
            s.ChannelMessageSend(m.ChannelID, "Error starting story: "+err.Error())
            return
        }
        s.ChannelMessageSend(m.ChannelID, "Story started! Theme: "+theme+", Tone: "+tone+". Add with !add <text>")
    } else if strings.HasPrefix(content, "!add ") {
        contribution := strings.TrimPrefix(content, "!add ")
        story, err := db.GetActiveStory(m.ChannelID)
        if err != nil || story == nil {
            s.ChannelMessageSend(m.ChannelID, "No active story! Start one with !story start")
            return
        }
        if !sentiment.MatchesTone(contribution, story.Tone) {
            s.ChannelMessageSend(m.ChannelID, "Contribution doesn't match story tone ("+story.Tone+")")
            return
        }
        err = db.AddContribution(m.ChannelID, m.Author.ID, contribution)
        if err != nil {
            s.ChannelMessageSend(m.ChannelID, "Error adding contribution: "+err.Error())
            return
        }
        s.ChannelMessageSend(m.ChannelID, "Added: "+contribution)
    } else if strings.HasPrefix(content, "!branch ") {
        branchText := strings.TrimPrefix(content, "!branch ")
        msg, err := s.ChannelMessageSend(m.ChannelID, "Vote for branch: "+branchText+" (react with 👍/👎)")
        if err != nil {
            s.ChannelMessageSend(m.ChannelID, "Error proposing branch: "+err.Error())
            return
        }
        db.SaveBranchProposal(msg.ID, m.ChannelID, branchText, m.Author.ID)
        s.MessageReactionAdd(m.ChannelID, msg.ID, "👍")
        s.MessageReactionAdd(m.ChannelID, msg.ID, "👎")
    }
}
