package main

import (
	"log"
	"net/http"

	"github.com/oldtimeguitarguy/todo/todo"
)

type MarshalJSON func(v interface{}) ([]byte, error)

type todoWriter struct {
	MarshalJSON MarshalJSON
	Res         http.ResponseWriter
}

func (w todoWriter) Write(todo todo.Todo, status todo.Status) {
	w.Res.Header().Set("Content-Type", "application/json")

	jsonData, err := w.MarshalJSON(todo)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		w.Res.WriteHeader(500)
		jsonData = []byte("{\"ERROR\":500}")
	}

	w.Res.Write(jsonData)
}
