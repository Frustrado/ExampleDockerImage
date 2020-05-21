package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TestMsg struct {
	Test string `json:"test"`
}

func showExample(w http.ResponseWriter, r *http.Request) {
	data := "Test message"
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func createExample(w http.ResponseWriter, r *http.Request) {
	var t TestMsg
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(t.Test)
}
