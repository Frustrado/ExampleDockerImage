package main

import (
	"github.com/bmizerany/pat"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func (app *Application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(showExample))
	mux.Get("/all", http.HandlerFunc(app.showAllExamples))
	mux.Post("/", http.HandlerFunc(app.createExample))
	mux.Get("/metrics", promhttp.Handler())
	return mux
}
