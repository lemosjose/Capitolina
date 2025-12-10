package download

import (
	"errors"
	"os"
	"strings"

	"github.com/lemosjose/capitu/bot/handlers/messages"
	"github.com/lemosjose/capitu/features/apis/googlebooks"

	telego "github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func formatMessage(sb *strings.Builder, header string, list []string) {
	sb.WriteString(header + "\n")
	for _, link := range list {
		sb.WriteString(link)
		sb.WriteString("\n \n")
	}
	sb.WriteString("\n")
}

// TODO (as soon as infrastructure helps) - return a file directly together with the links
func Downloadhandler(ctx *th.Context, update telego.Update) error {
	_, _, args := tu.ParseCommand(update.Message.Text)

	if len(args) == 0 {
		messages.SendError(ctx, update, errors.New("Please provide at least the book title"))
		return nil
	}

	target := "(accessInfo(epub(acsTokenLink),pdf(acsTokenLink)))"
	key := os.Getenv("BOOK_API_KEY")

	book := strings.Join(args, " ")
	author := ""

	if len(args) > 1 {
		author = args[1]
	}

	_, pdfs, epubs, err := googlebooks.GetBook(book, author, target, key)
	if err != nil {
		messages.SendError(ctx, update, errors.New("Could not get a download link from any of the sources"))
		return nil
	}

	var sb strings.Builder

	if len(pdfs) > 0 {
		formatMessage(&sb, "PDFs", pdfs)
	} else if len(epubs) > 0 {
		formatMessage(&sb, "EPUBs", epubs)
	}

	if sb.Len() == 0 {
		messages.SendError(ctx, update, errors.New("No download links found"))
		return nil
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		sb.String(),
	).WithParseMode(telego.ModeMarkdown))

	return nil
}
