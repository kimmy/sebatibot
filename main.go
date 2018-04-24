package main

import (
	"os"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot"

	"sebatibot/logger"
	"sebatibot/routes"
)

func main() {
	log := logger.GetLogger()
	app := mux.NewRouter()
	addRoutes(app, routes.UsersRoutes, "/")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	telegramToken := os.Getenv("TOKEN")
	log.Debugf("TOKEN: %s", telegramToken)

	// Create new telegram bot server using token
	withWebhook := tbot.WithWebhook("https://4f328400.ngrok.io/updates", "0.0.0.0:3000")
	bot, err := tbot.NewServer(telegramToken, withWebhook)
	if err != nil {
		log.Fatal(err)
	}

	go bot.ListenAndServe()
	log.Info("Listening at 0.0.0.0:3000 for telegram updates...")
	http.ListenAndServe("0.0.0.0:3000", app)
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
