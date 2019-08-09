package service

import "github.com/oldtimeguitarguy/todo/app/command/service"

type Todos interface {
	Create(label string) (*service.TodoModel, error)
	UpdateByID(id interface{}, label string, isComplete bool) (TodoModel, error)
	DeleteByID(id interface{}) error
	GetByID(id interface{}) (TodoModel, error)
	GetList() (TodoListModel, error)
}

type Logger interface {
	LogError(error)
}

type TodoListModel interface {
	GetLength() int
	ForEach(func(todo TodoModel, index int))
}
