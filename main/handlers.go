package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type test_msg struct {
	Test string
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var t test_msg
	err = json.Unmarshal(body, t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Test)
}