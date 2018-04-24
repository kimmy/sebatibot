package controllers

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/yanzay/tbot"

	"sebatibot/logger"
	"sebatibot/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Updates(w http.ResponseWriter, r *http.Request) {
	var telegramUpdate models.TelegramUpdate
	 log := logger.GetLogger()

	// Get TelegramUpdate Params
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bodyBytes, &telegramUpdate); err != nil {
 		log.Fatal(err)
 	}

 	log.WithFields(logrus.Fields{
		"telegram_update": telegramUpdate,
 	}).Info("Telegram Update Webhook")

 	// Initialize bot library
	bot, _ := tbot.NewServer(os.Getenv("TOKEN"))
	bot.Send(telegramUpdate.Message.Chat.Id, "Hi there!")

	// With Webhooks, you should do your own message parsing
  // and create necessary handlers for them.
  // 
  // With Polling, library provides it.
}