package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func createTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tw := todoWriter{Res: w, MarshalJSON: json.Marshal}
	todosController.CreateTodo(tw, r.FormValue("label"))
}
