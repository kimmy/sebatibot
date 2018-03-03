package main

import (
	// "fmt"
	"net/http"
	// "os"

	// "github.com/markbates/going/defaults"
	"github.com/go-telegram-bot-api/telegram-bot-api"

	"sebatibot/app"
	"sebatibot/logger"
	"github.com/joho/godotenv"
)

func main() {

	app := app.App()
	log := logger.GetLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := defaults.String(os.Getenv("GO_ENV_PORT"), "8000")

	log.Debugf("Starting API server at %s", port)
	log.Debugf("TOKEN: %s", os.Getenv("TOKEN"))

	//Start the api here
	// http.ListenAndServe(fmt.Sprintf(":%s", port), app)

	// Initializing Telegram Bot . . .


	// // ---
	// bot, err := tgbotapi.NewBotAPI("551123175:AAGsVKDZTiQ4Bz4WBHSgP43r7nzATeCAn3U")
	// if err != nil {
	// 	log.Panic(err)
	// }

	// bot.Debug = true

	// log.Debugf("Authorized on account %s", bot.Self.UserName)

	// _, err = bot.SetWebhook(tgbotapi.NewWebhook("https://d073e447.ngrok.io/551123175:AAGsVKDZTiQ4Bz4WBHSgP43r7nzATeCAn3U"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// updates := bot.ListenForWebhook("/551123175:AAGsVKDZTiQ4Bz4WBHSgP43r7nzATeCAn3U")
	// // ---


	go http.ListenAndServe("127.0.0.1:3000", app)
	// go http.ListenAndServe("127.0.0.1:8443")
	log.Debugf("Listening at 127.0.0.1:3000 for Telegram updates. . .")


	for update := range updates {
		log.Debugf("%+v\n", update)
	}
	
}
