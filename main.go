package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Router
	router := httprouter.New()

	router.POST("/todos", createTodoHandler)

	// Server
	log.Fatal(http.ListenAndServe(":8080", router))
}
