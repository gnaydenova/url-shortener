package main

import (
	"context"
	"net/http"

	"github.com/gnaydenova/url-shortener/app/handlers"
	"github.com/gnaydenova/url-shortener/app/shortener"
	"github.com/gnaydenova/url-shortener/app/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	repo := storage.NewMongoURLRepository(client.Database("url_shortener").Collection("urls"))

	http.Handle("/", handlers.NewURLHandler(shortener.NewShortener(repo)))

	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}
