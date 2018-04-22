package main

import (
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gorilla/mux"

	"sebatibot/logger"
	"github.com/joho/godotenv"
	"sebatibot/routes"
)

func main() {

	app := mux.NewRouter()
	addRoutes(app, routes.UsersRoutes, "/")
	log := logger.GetLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	telegramToken := os.Getenv("TOKEN")
	log.Debugf("TOKEN: %s", telegramToken)

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal("Error initializing tgbotapi")
	}

	bot.Debug = true
	log.Info("Authorized on account: ", bot.Self.UserName)

	// ngrok for local development
	// Run ngrok on your local and change webhookUrl by the generated one
	// Consult README.md for more details

	// webhookUrl := "https://{production_api_endpoint}/" + os.Getenv("TOKEN")
	webhookUrl := "https://a7685903.ngrok.io/" + bot.Token

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookUrl))
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Listening at 127.0.0.1:3000 for Telegram updates. . .")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", app))
}

func addRoutes(app *mux.Router, routes routes.Routes, prefix string) {

	for _, route := range routes {

		if prefix == "" {
			app.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				HandlerFunc(route.HandlerFunc)
		} else {
			app.
				PathPrefix(prefix).
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				HandlerFunc(route.HandlerFunc)
		}
	}
}
