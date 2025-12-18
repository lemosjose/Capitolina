package book

import (
	"fmt"

	"github.com/lemosjose/capitu/bot/handlers/messages"
	"github.com/lemosjose/capitu/features/apis/openlibrary"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func OpenLibraryHandler(ctx *th.Context, update telego.Update) error {
	_, _, args := tu.ParseCommand(update.Message.Text)
	if len(args) < 2 {
		messages.SendError(ctx, update, fmt.Errorf("Please provide both the book title and author"))
		return nil
	}

	title := args[0]
	author := args[1]

	url, err := openlibrary.GetBook(title, author)
	if err != nil {
		messages.SendError(ctx, update, err)
		return nil
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		url,
	).WithParseMode(telego.ModeMarkdown))

	return nil
}
