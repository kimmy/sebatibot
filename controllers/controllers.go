package controllers

import (
	"net/http"
	"io/ioutil"
	"fmt"

	"sebatibot/logger"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TelegramUpdate(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	log.Debugf("Request Body: %s", bodyString)
	w.WriteHeader(http.StatusOK)
}