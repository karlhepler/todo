package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func createTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todosController.CreateTodo(todoWriter{w}, r.FormValue("label"))
}

type Todo struct {
	ID         interface{}
	Label      string
	IsComplete bool
}

func showTodoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	todosController.GetTodoByID(todoWriter{w}, ps.ByName("id"))
}
