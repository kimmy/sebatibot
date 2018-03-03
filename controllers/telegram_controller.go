package controllers

import (
  // "encoding/json"
  "net/http"
  "io/ioutil"

  // "github.com/Sirupsen/logrus"

  "sebatibot/logger"
  // "sebatibot/models"
)

func init() {
  log = logger.GetLogger()
}

// UsersController defines the http routes for User Resource
type TelegramController struct {
  BaseController
}

// Post is used to create a new user object
func (u *TelegramController) Post(w http.ResponseWriter, r *http.Request) {

  // var user models.User
  // decoder := json.NewDecoder(r.Body)
  // err := decoder.Decode(&user)

  // if err != nil {
  //   log.Error(err)
  //   w.WriteHeader(http.StatusBadRequest)
  //   return

  // }

  bodyBytes, _ := ioutil.ReadAll(r.Body)
  bodyString := string(bodyBytes)

  log.Debugf("Request Body: %s", bodyString)

  // Write to db here

  w.WriteHeader(http.StatusOK)
  return
}
