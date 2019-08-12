package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oldtimeguitarguy/todo/todo"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB
var todosController *todo.TodosController

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
	todosController = &todo.TodosController{
		Logger: &MyLogger{
			Log: log.New(os.Stderr, "", log.LstdFlags),
		},
		TodoRepository: &SQLiteDriver{DB: db},
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
