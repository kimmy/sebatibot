package controllers

import (
  // "encoding/json"
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

  // Write to db here
  w.WriteHeader(http.StatusOK)
}

// // Post is used to create a new user object
// func (u *TelegramController) Post(w http.ResponseWriter, r *http.Request) {

//   // var user models.User
//   // decoder := json.NewDecoder(r.Body)
//   // err := decoder.Decode(&user)

//   // if err != nil {
//   //   log.Error(err)
//   //   w.WriteHeader(http.StatusBadRequest)
//   //   return

//   // }
//   bodyBytes, _ := ioutil.ReadAll(r.Body)
//   bodyString := string(bodyBytes)

//   log.Debugf("Request Body: %s", bodyString)

//   // Write to db here

//   w.WriteHeader(http.StatusOK)
//   fmt.Fprint(w, "Welcome!\n")
//   return
// }
