package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *Application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(showExample))
	mux.Get("/all", http.HandlerFunc(app.showAllExamples))
	mux.Post("/", http.HandlerFunc(app.createExample))
	return mux
}
