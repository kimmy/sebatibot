package main

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot"

	"sebatibot/logger"
	"sebatibot/routes"
	"sebatibot/controllers"
)

func main() {
	log := logger.GetLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	telegramToken := os.Getenv("TOKEN")
	log.Debugf("TOKEN: %s", telegramToken)

	// Create new telegram bot server using token
	bot, err := tbot.NewServer(telegramToken)
	if err != nil {
		log.Fatal(err)
	}

	// Handle with HiHandler function
	// Check library for other possible stuff: https://github.com/yanzay/tbot
	bot.HandleFunc("/hi", controllers.HiHandler)

	// You can also do this directly
	bot.Handle("hi", "test")

	log.Info("Telegram bot starting. . .")
	// Start listening for messages
	err = bot.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
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
