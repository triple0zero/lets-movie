package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/triple0zero/lets-movie/internal/botkit"
	"github.com/triple0zero/lets-movie/internal/model"
	"math"
)

type MovieListStorage interface {
	Movies(ctx context.Context) ([]model.Movie, error)
}

func ViewCmdList(list MovieListStorage) botkit.ViewFunc {

	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

		movies, err := list.Movies(ctx)
		if err != nil {
			return err
		}

		t := "Список фильмов:\n\n"

		for _, movie := range movies {
			t += fmt.Sprintf("%d: %s; pk: %f, imdb: %f > %s\n", movie.ID, movie.Name, math.Round(100*movie.KpRating)/100, math.Round(100*movie.ImdbRating)/100, movie.CreatedAt.Format("YYYY-MM-DD"))
		}

		var (
			msgText = t
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
