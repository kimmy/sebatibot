package controllers

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"

	"github.com/Sirupsen/logrus"
	"github.com/go-telegram-bot-api/telegram-bot-api"

	"sebatibot/logger"
	"sebatibot/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TelegramUpdate(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	var telegramUpdate models.TelegramUpdate
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bodyBytes, &telegramUpdate); err != nil {
		log.Fatal(err)
	}

	log.WithFields(logrus.Fields{
		"telegram_update": telegramUpdate,
	}).Info("Telegram Update Webhook")

	telegramToken := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal("Error initializing tgbotapi")
	}

	// Uncomment this if you want to see additional logging from `tgbotapi` library
	// bot.Debug = true
	// log.Info("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(telegramUpdate.Message.Chat.Id, "This is a sample reply.")
	// Uncoment this if you want to send a message as a `reply`
	// msg.ReplyToMessageID = telegramUpdate.Message.MessageId

	bot.Send(msg)

	w.WriteHeader(http.StatusOK)
}