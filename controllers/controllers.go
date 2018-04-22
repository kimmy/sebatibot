package controllers

import (
	"net/http"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/yanzay/tbot"

	"sebatibot/logger"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func HiHandler(message *tbot.Message) {
	log := logger.GetLogger()

	log.WithFields(logrus.Fields{
		"telegram_update": message.Text(),
	}).Info("HiHandler")

	// Handler can reply with several messages
	// Check model for possible fields:
	// https://github.com/yanzay/tbot/blob/4cc12770de420a77015159bb66f8a7173a2e4c1b/model/message.go
	message.Replyf("Hello %s", message.From.FirstName)
	time.Sleep(1 * time.Second)
	message.Reply("What's up?")
}