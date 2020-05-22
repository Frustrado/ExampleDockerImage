package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

type ServerConfig struct {
	Addr  string
	DbUri string
}

type Application struct {
	Data *AppData
}

func main() {
	cfg := ServerConfig{
		Addr:  getEnv("test_addr", ":8080"),
		DbUri: "mongodb://" + getEnv("MY_APP_DB_HOST", "database.default.svc.cluster.local"),
	}

	dbClient, err := cfg.openDB()
	if err != nil {
		log.Print(cfg.DbUri)
		log.Fatal(err)
	}
	dbData := &AppData{DB: dbClient}
	defer dbClient.Disconnect(context.TODO())
	log.Print("DB connection opened")

	app := &Application{
		Data: dbData,
	}

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: app.routes(),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

func (cfg *ServerConfig) openDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.DbUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
