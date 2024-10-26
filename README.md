# Go Rem Bot

Go Rem is a Discord bot built in Go that interacts with users in a unique and introspective way. The bot utilizes the Ollama API to generate responses based on user prompts, providing a blend of reality and fiction in its interactions.

[Readme in Portuguese](README_PT_BR.md)

[![en](https://img.shields.io/badge/lang-en-red.svg)](README.md)
[![pt-br](https://img.shields.io/badge/lang-pt--br-green.svg)](README_PT_BR.md)

## Features

- Responds to user commands in Discord channels.
- Maintains a history of conversations to provide contextually relevant responses.
- Uses a unique prompt template to guide its responses.
- Configurable via environment variables.

## Installation

### Prerequisites

- Go 1.23.2 or higher
- A Discord bot token
- Ollama API running locally

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/gabehamasaki/go-rem.git
   cd go-rem
   ```

2. Create a `.env` file in the root directory with the following content:

   ```env
   BOT_TOKEN=your_discord_bot_token
   BOT_PREFIX=!
   BOT_CHAT_ID=your_chat_id
   ```

3. Install the required Go modules:

   ```bash
   go mod tidy
   ```

4. Run the bot:

   ```bash
   make run
   ```

## Usage

- The bot responds to commands prefixed with the specified `BOT_PREFIX`.
- The default command is `talk`, which allows users to engage in conversation with the bot.
- The bot also has a `world` command that responds with a simple greeting.

## Commands

- `!talk <message>`: Engage in a conversation with the bot.
- `!world`: Receive a "Hello World!" message.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [DiscordGo](https://github.com/bwmarrin/discordgo) for the Discord API wrapper.
- [Ollama](https://ollama.com/) for the AI model used in generating responses.
