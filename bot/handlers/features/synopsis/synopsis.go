package synopsis

import (
	"errors"
	"os"
	"strings"

	"github.com/lemosjose/capitu/bot/handlers/messages"
	"github.com/lemosjose/capitu/features/apis/googlebooks"
	"github.com/lemosjose/capitu/features/genai"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func AiSynopsisHandler(ctx *th.Context, update telego.Update) error {
	_, _, args := tu.ParseCommand(update.Message.Text)
	if len(args) == 0 {
		messages.SendError(ctx, update, errors.New("Please provide at least the book title"))
		return nil
	}

	book := strings.Join(args, " ")
	synopsis, _, _, err := googlebooks.GetBook(book, "", "(volumeInfo(description))", os.Getenv("BOOK_API_KEY"))
	if err != nil {
		messages.SendError(ctx, update, errors.New("AI could not generate a synopsis"))
		return nil
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		string(synopsis),
	).WithParseMode(telego.ModeMarkdown))

	return nil
}

func SynopsisHandler(ctx *th.Context, update telego.Update) error {
	_, _, args := tu.ParseCommand(update.Message.Text)
	if len(args) == 0 {
		messages.SendError(ctx, update, errors.New("Please provide a book title"))
		return nil
	}

	key := os.Getenv("BOOK_API_KEY")
	target := "(volumeInfo(description))"

	book := strings.Join(args, " ")
	author := ""

	if len(args) > 1 {
		author = args[1]
	}

	description, _, _, err := googlebooks.GetBook(book, author, target, key)

	if err == nil && len(description) > 0 {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			string(description),
		).WithParseMode(telego.ModeMarkdown))
		return nil
	}

	synopsis, aiErr := genai.GeminiFetch(book, author)
	if aiErr != nil {
		messages.SendError(ctx, update, errors.New("Could not find a synopsis (Google Books empty, AI failed)"))
		return nil
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		synopsis,
	).WithParseMode(telego.ModeMarkdown))

	return nil
}
