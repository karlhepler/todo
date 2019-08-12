package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func createTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todosController.CreateTodo(todoWriter{w}, r.FormValue("label"))
}
