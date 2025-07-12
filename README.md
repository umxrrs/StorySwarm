# StorySwarm
---

**StorySwarm** is a Discord bot written in Go that enables collaborative storytelling in Discord servers. Users can start a story with a theme and tone, contribute sentences, propose branching plotlines with votes, and export stories as HTML pages. The bot uses sentiment analysis to ensure contributions match the story's tone.

## Features

- **Collaborative Storytelling**: Start a story with `!story start <theme> <tone>` and add contributions with `!add <text>`.
- **Story Branching**: Propose alternate plotlines with `!branch <text>` and vote using reactions.
- **Tone Consistency**: Simple keyword-based sentiment analysis ensures contributions match the story's tone (e.g., epic, spooky).
- **Web Export**: View completed stories at `http://<server>:8080/stories/<channel-id>`.
---

## Prerequisites

- **Go**: Version 1.20 or higher
- **SQLite**: For story storage
- **Discord Bot Token**: Obtain from [Discord Developer Portal](https://discord.com/developers/applications)
---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/umxrrs/StorySwarm.git
   cd StorySwarm
   ```
---

2. Install dependencies:
   ```bash
   go mod tidy
   ```
---

3. Create a config.json file in the config directory: (already exists though)
    ```json
    {
    "token": "YOUR_DISCORD_BOT_TOKEN",
    "web_port": "8080"
}
    ```

---

4. Build and run:
   ```bash
   go build -o storyswarm cmd/main.go
   ./storyswarm
   ```
## Usage
1. Invite the bot to your Discord server using the link from the Discord Developer Portal.
2. Use commands in any channel:
- `!story start fantasy epic`: Start a fantasy story with an epic tone.
- `!add The hero ventured into the dark forest.`: Add a contribution.
- `!branch The hero finds a hidden castle`.: Propose a story branch (vote with 👍/👎).
View the story at `http://<server>:8080/stories/<channel-id>.`

---

## Testing
Run unit tests:
 ```bash
go test ./...
```
---

## Hosting

**Local**: Run `./storyswarm` on your machine.
**Cloud**: Deploy on `Heroku`, `AWS`, or `DigitalOcean`. Ensure port 8080 is open for the web server.

---

# Contributing

---

Contributions are welcome! Please submit a pull request or open an issue on [GitHub.](https://github.com/umxrrs/StorySwarm)

---

# License
This project is licensed under the MIT License. See the LICENSE file for details.

---

# Author
- Created by [umxrrs/umar](https://github.com/umxrrs)
