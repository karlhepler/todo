package main

import (
	"database/sql"
	"log"
)

type Status uint8

const (
	StatusOK Status = iota
	StatusError
	StatusCreated
)

type Todo struct {
	ID         interface{} `json:"id"`
	Label      string      `json:"label"`
	IsComplete bool        `json:"isComplete"`
}

type TodoWriter interface {
	Write(Todo, Status)
}

type TodosController struct {
	Log *log.Logger
	DB  *sql.DB
}

func (ctrl TodosController) CreateTodo(w TodoWriter, label string) {
	// Make the todo
	todo := Todo{
		Label:      label,
		IsComplete: false,
	}

	tx, err := ctrl.DB.Begin()
	if err != nil {
		ctrl.Log.Print("ERROR: " + err.Error())
		w.Write(todo, StatusError)
		return
	}
	stmt, err := tx.Prepare("insert into todos(label, is_complete) values(?,?)")
	if err != nil {
		ctrl.Log.Print("ERROR: " + err.Error())
		w.Write(todo, StatusError)
		return
	}
	defer stmt.Close()

	rowsAffected, err := stmt.Exec(todo.Label, todo.IsComplete)
	if err != nil {
		ctrl.Log.Print("ERROR: " + err.Error())
		w.Write(todo, StatusError)
		return
	}

	id, err := rowsAffected.LastInsertId()
	if err != nil {
		ctrl.Log.Print("ERROR: " + err.Error())
		w.Write(todo, StatusError)
		return
	}

	tx.Commit()

	// Set the id
	todo.ID = id

	w.Write(todo, StatusCreated)
}
