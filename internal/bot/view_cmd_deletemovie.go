package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/triple0zero/lets-movie/internal/botkit"
	"github.com/triple0zero/lets-movie/internal/model"
)

type MovieDeleter interface {
	Delete(ctx context.Context, id int64) error
}

func ViewCmdRemoveMovie(deleter MovieDeleter) botkit.ViewFunc {

	type DeleteMovieArgs struct {
		ID int64 `json:"id"`
	}

	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

		args, err := botkit.ParseJSON[DeleteMovieArgs](update.Message.CommandArguments())
		if err != nil {
			return err
		} //TODO: Описать ошибку невернного инпута

		movie := model.Movie{
			ID: args.ID,
		}

		err = deleter.Delete(ctx, movie.ID)
		if err != nil {
			return err
		}

		var (
			msgText = fmt.Sprintf("Фильм с ID: %d удален.", movie.ID)
			reply   = tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		)

		//reply.ParseMode = "ModeMarkdownV2"

		_, err = bot.Send(reply)
		if err != nil {
			return err
		}
		return nil
	}
}
