package main

import (
	"minimalgo/api"
	"minimalgo/utils"

	log "github.com/sirupsen/logrus"
)

func init() {
	utils.LogSettting()
}

func main() {

	app := api.Setup()

	// Start server
	log.Fatal(app.Listen(5430))

}
