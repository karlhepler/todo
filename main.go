package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB
var todosController *TodosController

func main() {
	// Database
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the todos table
	sqlStmt := `create table todos (id integer not null primary key, label text not null, is_complete integer not null);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	// setup controller
	todosController = &TodosController{
		Log: log.New(os.Stderr, "", log.LstdFlags),
		DB:  db,
	}

	// Router
	router := httprouter.New()

	router.POST("/todos", createTodoHandler)

	// Server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tw := todoWriter{Res: w, MarshalJSON: json.Marshal}
	todosController.CreateTodo(tw, r.FormValue("label"))
}

type MarshalJSON func(v interface{}) ([]byte, error)

type todoWriter struct {
	MarshalJSON MarshalJSON
	Res         http.ResponseWriter
}

func (w todoWriter) Write(todo Todo, status Status) {
	w.Res.Header().Set("Content-Type", "application/json")

	jsonData, err := w.MarshalJSON(todo)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		w.Res.WriteHeader(500)
		jsonData = []byte("{\"ERROR\":500}")
	}

	w.Res.Write(jsonData)
}
