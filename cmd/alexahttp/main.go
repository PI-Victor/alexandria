package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"

	"github.com/PI-Victor/alexandria/pkg/server"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	app := server.App{}

	http.HandleFunc("/", app.Index)

	logrus.Info("Server started...")
	http.ListenAndServe(":8080", nil)
}

func getRegistryManifests() error {
	return nil
}
