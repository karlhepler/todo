package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oldtimeguitarguy/todo/todo"
)

type todoWriter struct {
	Res http.ResponseWriter
}

func (w todoWriter) Write(todo todo.Todo, status todo.Status) {
	w.Res.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		w.Res.WriteHeader(500)
		jsonData = []byte("{\"ERROR\":500}")
	}

	w.Res.Write(jsonData)
}
