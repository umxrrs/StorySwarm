package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Story struct {
    ChannelID string
    Theme     string
    Tone      string
}

type Contribution struct {
    UserID string
    Text   string
}

type BranchProposal struct {
    MessageID string
    ChannelID string
    Text      string
    UserID    string
}

type Vote struct {
    MessageID string
    UserID    string
    Upvote    bool
}

var db *sql.DB

func Init() error {
    var err error
    db, err = sql.Open("sqlite3", "./stories.db")
    if err != nil {
        return err
    }
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS stories (
            channel_id TEXT PRIMARY KEY,
            theme TEXT,
            tone TEXT
        );
        CREATE TABLE IF NOT EXISTS contributions (
            channel_id TEXT,
            user_id TEXT,
            text TEXT
        );
        CREATE TABLE IF NOT EXISTS branch_proposals (
            message_id TEXT PRIMARY KEY,
            channel_id TEXT,
            text TEXT,
            user_id TEXT
        );
        CREATE TABLE IF NOT EXISTS votes (
            message_id TEXT,
            user_id TEXT,
            upvote INTEGER
        );
    `)
    return err
}

func SaveStory(channelID, theme, tone string) error {
    _, err := db.Exec("INSERT INTO stories (channel_id, theme, tone) VALUES (?, ?, ?)", channelID, theme, tone)
    return err
}

func GetActiveStory(channelID string) (*Story, error) {
    row := db.QueryRow("SELECT channel_id, theme, tone FROM stories WHERE channel_id = ?", channelID)
    var story Story
    err := row.Scan(&story.ChannelID, &story.Theme, &story.Tone)
    if err != nil {
        return nil, err
    }
    return &story, nil
}

func AddContribution(channelID, userID, text string) error {
    _, err := db.Exec("INSERT INTO contributions (channel_id, user_id, text) VALUES (?, ?, ?)", channelID, userID, text)
    return err
}

func GetContributions(channelID string) ([]Contribution, error) {
    rows, err := db.Query("SELECT user_id, text FROM contributions WHERE channel_id = ?", channelID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var contributions []Contribution
    for rows.Next() {
        var c Contribution
        rows.Scan(&c.UserID, &c.Text)
        contributions = append(contributions, c)
    }
    return contributions, nil
}

func SaveBranchProposal(messageID, channelID, text, userID string) error {
    _, err := db.Exec("INSERT INTO branch_proposals (message_id, channel_id, text, user_id) VALUES (?, ?, ?, ?)", messageID, channelID, text, userID)
    return err
}

func GetBranchProposal(messageID string) (*BranchProposal, error) {
    row := db.QueryRow("SELECT message_id, channel_id, text, user_id FROM branch_proposals WHERE message_id = ?", messageID)
    var p BranchProposal
    err := row.Scan(&p.MessageID, &p.ChannelID, &p.Text, &p.UserID)
    if err != nil {
        return nil, err
    }
    return &p, nil
}

func AddVote(messageID, userID string, upvote bool) error {
    _, err := db.Exec("INSERT INTO votes (message_id, user_id, upvote) VALUES (?, ?, ?)", messageID, userID, upvote)
    return err
}

func GetVoteCount(messageID string) (upvotes, downvotes int, err error) {
    rows, err := db.Query("SELECT upvote FROM votes WHERE message_id = ?", messageID)
    if err != nil {
        return 0, 0, err
    }
    defer rows.Close()
    for rows.Next() {
        var upvote bool
        rows.Scan(&upvote)
        if upvote {
            upvotes++
        } else {
            downvotes++
        }
    }
    return upvotes, downvotes, nil
}

func ClearBranchProposal(messageID string) error {
    _, err := db.Exec("DELETE FROM branch_proposals WHERE message_id = ?", messageID)
    if err != nil {
        return err
    }
    _, err = db.Exec("DELETE FROM votes WHERE message_id = ?", messageID)
    return err
}
