package main

import (
	"log"
	"net/http"
	"os"
)

type ServerConfig struct {
	Addr  string
	DbUri string
}

type Application struct {
	Data []string
}

func main() {
	cfg := ServerConfig{
		Addr:  getEnv("test_addr", ":8080"),
		DbUri: getEnv("db_addr", "test"),
	}

	app := &Application{
		Data: make([]string, 0),
	}

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: app.routes(),
	}
	err := server.ListenAndServe()
	log.Fatal(err)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
