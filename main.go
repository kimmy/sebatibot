package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/markbates/going/defaults"
	"sebatibot/app"
	"sebatibot/logger"
)

func main() {

	app := app.App()
	port := defaults.String(os.Getenv("GO_ENV_PORT"), "8000")

	log := logger.GetLogger()
	log.Debugf("Starting API server at %s", port)

	//Start the api here
	http.ListenAndServe(fmt.Sprintf(":%s", port), app)
}
