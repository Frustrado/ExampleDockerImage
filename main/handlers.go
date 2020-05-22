package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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
	data, err := app.Data.Get()
	if err != nil {
		log.Fatal(err)
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
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
	log.Print("Inserting")
	insertResult, err := app.Data.Insert(t.Test)
	if err != nil {
		log.Print("Cannot insert")
		log.Fatal(err)
	}
	log.Print("Inserted")
	if _, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		http.Redirect(w, r, fmt.Sprintf("/all"), http.StatusSeeOther)
	}
}
