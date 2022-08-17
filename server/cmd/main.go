package main

import (
	"backViewer/internal/config"
	"backViewer/internal/mongo"
	"backViewer/pkg/handlers"
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

func main() {
	// Define env variables
	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(cfg)
	// Connection to mongoDB
	mr, err := mongo.NewConn(cfg.MongoDB)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mr.Client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	router := handlers.NewEmailHandler(mr)

	log.Fatal(http.ListenAndServe(cfg.Server.Port, router))
}
