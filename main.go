package main

import (
	"net/http"
	"os"
	"time"

	"github.com/hardsky/car-center-api/api"
	"github.com/hardsky/car-center-api/db"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)

	if len(os.Getenv("DP_DEBUG")) > 0 {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {

	d := db.Connect(&db.Opts{
		Addr:     os.Getenv("DP_DB_ADDR"),
		User:     os.Getenv("DP_DB_USER"),
		Password: os.Getenv("DP_DB_PSW"),
		Database: os.Getenv("DP_DB_DATABASE"),
	})
	log.Info("database is connected")

	a := api.NewAPI(&api.Opts{
		Db: d,
	})
	log.Info("api routes are established")

	srv := &http.Server{
		Addr:         os.Getenv("DP_ADDR"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Handler(),
	}

	log.Info("car-center-api is ready")

	log.Fatal(srv.ListenAndServe())
}
