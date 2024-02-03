package bot

import (
	"context"
	"github.com/triple0zero/lets-movie/internal/botkit"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/triple0zero/lets-movie/parameters"
)

func ViewCmdHelp() botkit.ViewFunc {

	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

		var (
			msgText = parameters.HelpMessage
			reply   = tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		)

		//reply.ParseMode = "ModeMarkdownV2"

		_, err := bot.Send(reply)
		if err != nil {
			return err
		}
		return nil
	}
}
