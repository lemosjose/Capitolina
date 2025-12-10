package main

import (
	"context"
	"fmt"
	"os"

	//this autoloads the .env (read joho/godotenv), do not find it strang that there is not a function loading the .env
	_ "github.com/joho/godotenv/autoload"
	//IMPORTANT: read google's documentation on restricting keys, that's why this source code enforces having one key
	//for the books api and another one for the genai service
	"github.com/lemosjose/capitu/bot/handlers/features/download"
	"github.com/lemosjose/capitu/bot/handlers/features/synopsis"
	"github.com/lemosjose/capitu/bot/handlers/messages"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	ctx := context.Background()

	botToken := os.Getenv("BOT_TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println("Failed to create bot:", err)
		os.Exit(1)
	}

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		fmt.Println("Failed to get updates, is telegram or your network down?", err)
	}

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		fmt.Println("Failed to create bot handler:", err)
		os.Exit(1)
	}

	bh.Handle(messages.StartHandler, th.CommandEqual("start"))
	bh.Handle(download.Downloadhandler, th.CommandEqual("download"))
	bh.Handle(synopsis.SynopsisHandler, th.CommandEqual("synopsis"))
	bh.Handle(synopsis.AiSynopsisHandler, th.CommandEqual("aisynopsis"))
	bh.Handle(messages.HelpHandler, th.CommandEqual("help"))

	err = bot.SetMyCommands(ctx, &telego.SetMyCommandsParams{
		Commands: []telego.BotCommand{
			{Command: "start", Description: "Start the bot and see welcome message"},
			{Command: "download", Description: "Get PDF/EPUB links: /download '<book>' '<author>' "},
			{Command: "synopsis", Description: "Get book summary: /synopsis '<book>' '<author>'"},
			{Command: "aisynopsis", Description: "Force AI summary: /aisynopsis '<book>' 'author'"},
			{Command: "help", Description: "Show a brief description of all commands"},
		},
	})

	defer bh.Stop()
	bh.Start()
}
