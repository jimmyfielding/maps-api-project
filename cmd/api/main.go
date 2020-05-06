package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type server struct {
	log        *logrus.Logger
	mapsClient googlemaps.IClient
}

var log = logrus.New()

func startServer() error {
	log.Out = os.Stdout
	viper.SetConfigFile("../../.secrets/apiconfig.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal("failed to parse config")
	}

	log.Infof("found config file: %s", viper.ConfigFileUsed())
	mapsAPIKey := viper.GetString("mapsAPIKey")
	if mapsAPIKey == "" {
		log.Fatal("../../.secrets/apiconfig.yaml is missing mapsAPIKey field")
	}

	mapsClient, err := googlemaps.NewClient(config.MapsAPIKey)
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal("failed to create google maps client")
	}

	s := &server{
		log:        log,
		mapsClient: mapsClient,
	}

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: s.routes(),
	}

	log.Info("starting server on 127.0.0.1:8080")
	return srv.ListenAndServe()
}