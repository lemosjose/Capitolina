package main

import (
	"context"
	"fmt"
	"os"

	//this autoloads the .env (read joho/godotenv), do not find it strang that there is not a function loading the .env
	_ "github.com/joho/godotenv/autoload"

	"github.com/lemosjose/capitu/bot/handlers/features/googleBooks/download"
	"github.com/lemosjose/capitu/bot/handlers/features/googleBooks/synopsis"
	"github.com/lemosjose/capitu/bot/handlers/features/openLibrary/book"
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
	bh.Handle(messages.AboutHandler, th.CommandEqual("about"))
	bh.Handle(book.OpenLibraryHandler, th.CommandEqual("openlibrary"))

	err = bot.SetMyCommands(ctx, &telego.SetMyCommandsParams{
		Commands: []telego.BotCommand{
			{Command: "start", Description: "Start the bot and see welcome message"},
			{Command: "download", Description: "Get PDF/EPUB links: /download '<book>' '<author>' "},
			{Command: "synopsis", Description: "Get book summary: /synopsis '<book>' '<author>'"},
			{Command: "aisynopsis", Description: "Force AI summary: /aisynopsis '<book>' 'author'"},
			{Command: "help", Description: "Show a brief description of all commands"},
			{Command: "about", Description: "Shows relevant information about bot and dev"},
			{Command: "openlibrary", Description: "Gets book link from OpenLibrary: /openlibrary '<book>' '<author>'"},
		},
	})

	defer bh.Stop()
	bh.Start()
}
