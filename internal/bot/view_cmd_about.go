package bot

import (
	"context"
	"github.com/triple0zero/lets-movie/internal/botkit"
	"github.com/triple0zero/lets-movie/parameters"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ViewCmdAbout() botkit.ViewFunc {

	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

		var (
			msgText = parameters.AboutMessage
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
