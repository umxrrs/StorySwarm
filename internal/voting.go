package bot

import (
    "github.com/umxrrs/StorySwarm/internal/db"
    "github.com/bwmarrin/discordgo"
)

func HandleVote(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
    proposal, err := db.GetBranchProposal(r.MessageID)
    if err != nil || proposal == nil {
        return
    }
    if r.Emoji.Name == "👍" {
        db.AddVote(r.MessageID, r.UserID, true)
    } else if r.Emoji.Name == "👎" {
        db.AddVote(r.MessageID, r.UserID, false)
    }
    upvotes, downvotes, _ := db.GetVoteCount(r.MessageID)
    if upvotes+downvotes >= 5 {
        if upvotes > downvotes {
            db.AddContribution(proposal.ChannelID, proposal.UserID, proposal.Text)
            s.ChannelMessageSend(r.ChannelID, "Branch accepted: "+proposal.Text)
        } else {
            s.ChannelMessageSend(r.ChannelID, "Branch rejected: "+proposal.Text)
        }
        db.ClearBranchProposal(r.MessageID)
    }
}
