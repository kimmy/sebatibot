package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/markbates/going/defaults"
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
	http.ListenAndServe(fmt.Sprintf(":%s", port), app)
}
