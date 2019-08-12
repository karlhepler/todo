package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/oldtimeguitarguy/todo/driver/logger"
	"github.com/oldtimeguitarguy/todo/driver/repo"
	"github.com/oldtimeguitarguy/todo/todo"
)

var db *sql.DB
var todosController *todo.TodosController

func init() {
	// Database
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	// Create the todos table
	sqlStmt := `create table todos (id integer not null primary key, label text not null, is_complete integer not null);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	// Init gateway drivers
	l := &logger.MyLogger{
		Log: log.New(os.Stderr, "", log.LstdFlags),
	}
	todoRepo := &repo.SQLiteDriver{DB: db}

	// setup controller
	todosController = todo.NewTodosController(l, todoRepo)
}
