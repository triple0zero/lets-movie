package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/triple0zero/lets-movie/internal/botkit"
	"github.com/triple0zero/lets-movie/internal/model"
	"math"
)

const (
	kinopoiskHost string = "kinopoisk.ru"
)

type MovieStorage interface {
	Add(ctx context.Context, movie model.Movie) (int64, error)
}

func ViewCmdAddMovieByName(storage MovieStorage) botkit.ViewFunc {

	type AddMovieArgs struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

		args, err := botkit.ParseJSON[AddMovieArgs](update.Message.CommandArguments())
		if err != nil {
			return err
		} //TODO: Описать ошибку невернного инпута

		movie := model.Movie{
			Name: args.Name,
			Url:  args.Url,
		}

		movieID, err := storage.Add(ctx, movie)
		if err != nil {
			return err
		}

		var (
			msgText = fmt.Sprintf("Фильм %s добавлен с ID: %d.", movie.Name, movieID)
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

func ViewCmdAddMovieByUrl(storage MovieStorage) botkit.ViewFunc {

	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

		movie := botkit.CreateMovieObj(update.Message.CommandArguments())

		movieID, err := storage.Add(ctx, movie)
		if err != nil {
			return err
		}

		var (
			msgText = fmt.Sprintf(
				`По указанной ссылке %s добавлен фильм %s 
						c рейтингами Кинопоиска: %f и IMDB: %f с ID: %d.\n\n
						Краткое описание: %s
`, movie.Url, movie.Name, math.Round(100*movie.KpRating)/100, math.Round(100*movie.ImdbRating)/100, movieID, movie.Description)
			reply = tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		)

		//reply.ParseMode = "ModeMarkdownV2"

		_, err = bot.Send(reply)
		if err != nil {
			return err
		}
		return nil
	}
}
