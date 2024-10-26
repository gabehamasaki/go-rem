# Go Rem Bot

Go Rem é um bot do Discord construído em Go que interage com os usuários de maneira única e introspectiva. O bot utiliza a API Ollama para gerar respostas com base em prompts dos usuários, proporcionando uma mistura de realidade e ficção em suas interações.

[Readme em Inglês](README.md)

[![en](https://img.shields.io/badge/lang-en-green.svg)](README.md)
[![pt-br](https://img.shields.io/badge/lang-pt--br-green.svg)](README_PT_BR.md)

## Features

- Responde a comandos dos usuários em canais do Discord.
- Mantém um histórico de conversas para fornecer respostas contextualmente relevantes.
- Utiliza um modelo de prompt único para guiar suas respostas.
- Configurável através de variáveis de ambiente.

## Installation

### Prerequisites

- Go 1.23.2 ou superior
- Um token de bot do Discord
- API Ollama rodando localmente

### Setup

1. Clone o repositório:

   ```bash
   git clone https://github.com/gabehamasaki/go-rem.git
   cd go-rem
   ```

2. Crie um arquivo `.env` no diretório raiz com o seguinte conteúdo:

   ```env
   BOT_TOKEN=seu_token_do_bot_discord
   BOT_PREFIX=!
   BOT_CHAT_ID=seu_chat_id
   ```

3. Instale os módulos Go necessários:

   ```bash
   go mod tidy
   ```

4. Execute o bot:

   ```bash
   make run
   ```

## Usage

- O bot responde a comandos prefixados com o `BOT_PREFIX` especificado.
- O comando padrão é `talk`, que permite aos usuários interagir em uma conversa com o bot.
- O bot também possui um comando `world` que responde com uma simples saudação.

## Commands

- `!talk <mensagem>`: Engaje em uma conversa com o bot.
- `!world`: Receba uma mensagem "Hello World!".

## Contributing

Contribuições são bem-vindas! Por favor, faça um fork do repositório e envie um pull request para quaisquer melhorias ou correções de bugs.

## License

Este projeto é licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## Acknowledgments

- [DiscordGo](https://github.com/bwmarrin/discordgo) pelo wrapper da API do Discord.
- [Ollama](https://ollama.com/) pelo modelo de IA usado na geração de respostas.
