package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	ID         interface{} `json:"id"`
	Label      string      `json:"label"`
	IsComplete bool        `json:"isComplete"`
}

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

	router.POST("/todos", createTodo)

	// Server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Make the todo
	todo := &Todo{
		Label:      r.FormValue("label"),
		IsComplete: false,
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into todos(label, is_complete) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rowsAffected, err := stmt.Exec(todo.Label, todo.IsComplete)
	if err != nil {
		log.Fatal(err)
	}

	id, err := rowsAffected.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	// Set the id
	todo.ID = id

	// Get the json data
	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
