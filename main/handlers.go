package main

import (
	"encoding/json"
	"net/http"
)

type TestMsg struct {
	Test string `json:"test"`
}

func showExample(w http.ResponseWriter, r *http.Request) {
	data := getEnv("MESSAGE_VALUE_KEY", "Hello world")
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (app *Application) showAllExamples(w http.ResponseWriter, r *http.Request) {
	data := app.Data
	dataJson, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataJson)
}

func (app *Application) createExample(w http.ResponseWriter, r *http.Request) {
	var t TestMsg
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	app.Data = append(app.Data, t.Test)
}
