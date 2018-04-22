package controllers

import (
	"net/http"
	// "net/url"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"

	"sebatibot/logger"
	"github.com/Sirupsen/logrus"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TelegramUpdate(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	var requestParams map[string]interface{}
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bodyBytes, &requestParams); err != nil {
		log.Fatal(err)
	}

	log.WithFields(logrus.Fields{
		"request_params": requestParams,
	}).Info("Telegram Update Webhook")

	telegramToken := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal("Error initializing tgbotapi")
	}

	bot.Debug = true
	// log.Info("Authorized on account %s", bot.Self.UserName)

	// msg := tgbotapi.NewMessage(-1001126778557, "Panget mu loko mo ko!")
	// msg := tgbotapi.NewMessage(253244466, "Panget mu loko mo ko!")
	// bot.Send(msg)

	// var data := TelegramApiUpdate{}

	// // Generate HTTP request URL and body
	// apiUrl := "https://api.telegram.org/bot" + os.Getenv("TOKEN") + "/sendMessage"
	// httpMethod := "POST"
	// data := url.Values{}
	// data.Add("chat_id", apiResponse.Message)
	// data.Add("text", "Hi there!")

	// // Generate actual HTTP request
	// req, err := http.NewRequest(httpMethod, apiUrl, strings.NewReader(data.Encode()))
	// if err != nil {
	// 	log.Fatal("HTTP error: ", err)
	// }

	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	// // Execute HTTP request
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatal("HTTP error: ", err)
	// }

	w.WriteHeader(http.StatusOK)
}