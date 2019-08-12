package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

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

	// Router
	router := httprouter.New()

	router.POST("/todos", createTodoHandler)

	// Server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tw := todoWriter{res: w}
	createTodo(tw, r.FormValue("label"))
}

type todoWriter struct {
	res http.ResponseWriter
}

func (w todoWriter) Write(todo Todo, status Status) {
	w.res.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		w.res.WriteHeader(500)
		jsonData = []byte("{\"ERROR\":500}")
	}

	w.res.Write(jsonData)
}
