package main

import (
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

type TodoRepository interface {
	Insert(label string, isComplete bool) (interface{}, error)
}

type TodosController struct {
	Log *log.Logger
	TodoRepository
}

func (ctrl TodosController) CreateTodo(w TodoWriter, label string) {
	// Make the todo
	todo := Todo{
		Label:      label,
		IsComplete: false,
	}

	id, err := ctrl.Insert(todo.Label, todo.IsComplete)
	if err != nil {
		w.Write(todo, StatusError)
		return
	}

	// Set the id
	todo.ID = id

	w.Write(todo, StatusCreated)
}
