package main

import (
	"backViewer/internal/mongo"
	"backViewer/pkg/handlers"
	"context"
	"flag"
	"log"
	"net/http"
)

func main() {

	// Define env variables
	var (
		host     = flag.String("host", ":8080", "The host of the application.")
		mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "Database connection URI")
	)

	flag.Parse()

	// Connection to mongoDB
	mr := mongo.NewConn(*mongoURI)
	defer func() {
		if err := mr.Client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	router := handlers.NewEmailHandler(mr)

	log.Fatal(http.ListenAndServe(*host, router))
}
