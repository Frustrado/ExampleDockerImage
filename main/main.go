package main

import (
	"log"
	"net/http"
	"os"
)
type ServerConfig struct {
	Addr string
	dbUri string
}
func main() {
	cfg := ServerConfig{
		Addr:  getEnv("test_addr", ":8080"),
		dbUri: getEnv("db_addr", "test"),
	}

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: routes(),
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

