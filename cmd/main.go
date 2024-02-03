package main

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//apiv1 "github.com/triple0zero/lets-movie/api/v1"
	"github.com/triple0zero/lets-movie/config"
	"github.com/triple0zero/lets-movie/internal/bot"
	"github.com/triple0zero/lets-movie/internal/botkit"
	"github.com/triple0zero/lets-movie/internal/storage"
	"go.uber.org/zap"
	"log"
)

const (
	telegramApiToken string = "LM_TELEGRAM_APITOKEN"
	dbDriverName     string = "postgres"
	dbUrl            string = "LM_DB_URL"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()

	botAPI, err := tgbotapi.NewBotAPI(config.GetEnvVariable(telegramApiToken))

	if err != nil {
		sugar.Error("Failed to create telegram bot api", zap.Error(err))
	}

	sugar.Infof("Authorized on account %s", botAPI.Self.UserName)

	//bot.Debug = true

	db, err := sqlx.Connect(dbDriverName, config.GetEnvVariable(dbUrl))
	if err != nil {
		sugar.Error("Failed to connect to db", zap.Error(err))
		return
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	sugar.Info("The application successfully connected to database")

	var (
		movieStorage = storage.NewMovieStorage(db)
	)

	lmBot := botkit.New(botAPI)

	lmBot.RegisterCmdView(
		"help",
		bot.ViewCmdHelp(),
	)

	lmBot.RegisterCmdView(
		"about",
		bot.ViewCmdAbout(),
	)

	lmBot.RegisterCmdView(
		"addByName",
		bot.ViewCmdAddMovieByName(movieStorage),
	)

	lmBot.RegisterCmdView(
		"add",
		bot.ViewCmdAddMovieByUrl(movieStorage),
	)

	lmBot.RegisterCmdView(
		"rm",
		bot.ViewCmdRemoveMovie(movieStorage),
	)

	lmBot.RegisterCmdView(
		"list",
		bot.ViewCmdList(movieStorage),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func(ctx context.Context) {
		if err := http.ListenAndServe("0.0.0.0:9874", mux); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Printf("[ERROR] failed to run http server: %v", err)
				return
			}

			log.Printf("[INFO] http server stopped")
		}
	}(ctx)

	if err := lmBot.Run(ctx); err != nil {
		sugar.Error("Failed to run botkit", zap.Error(err))
	}
}
