package main

import (
"github.com/bmizerany/pat"
"net/http"
)

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(showExample))
	mux.Post("/", http.HandlerFunc(createExample))
	return mux
}
