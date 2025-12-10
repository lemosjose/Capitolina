package messages

import (
	"fmt"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func StartHandler(ctx *th.Context, update telego.Update) error {
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Hello %s! I'm Capitu, i send download Links, synopsis on public domain books! type /help to see all the commands", update.Message.From.FirstName),
	))

	return nil
}

func AboutHandler(ctx *th.Context, update telego.Update) error {
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		"My programmer is José Lemos, reach him on linkedin.com/in/lemosjose and read his posts at lemosjose.github.io \n My source code is located in https://github.com/lemosjose/capitolina",
	))

	return nil
}

func HelpHandler(ctx *th.Context, update telego.Update) error {
	const helpMessage = `
	*/download <book> [author]*
	Searches for direct PDF or EPUB download links from Google Books. If none are found, it tries OpenLibrary.
	_Example: /download Hamlet Shakespeare_ 
	Prefer something like /download "Hamlet" "Shakespeare", keep that suggestion for other commands

	*/synopsis <book> [author]*
	Fetches the official book description from Google Books. If the description is missing, it automatically falls back to an AI-generated summary.
	_Example: /synopsis 1984 Orwell_

	*/aisynopsis <book> [author]*
	Forces the bot to generate a summary using Google Gemini 2.0 Flash, ignoring the official description. Great for shorter, creative summaries.
	_Example: /aisynopsis Dom Casmurro_

	*/about*
	Information about the creator and source code.

	*/help*
	Shows this message.
	`
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		helpMessage,
	).WithParseMode(telego.ModeMarkdown))

	return nil
}

func AnyHandler(ctx *th.Context, update telego.Update) error {
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		"This command does not exist. Send /help to see all commands",
	))

	return nil
}

func SendError(ctx *th.Context, update telego.Update, err error) {
	ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(update.Message.Chat.ID), err.Error()))
}
