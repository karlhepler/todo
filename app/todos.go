package app

import "github.com/oldtimeguitarguy/todo/app/services/todo"

type Status uint8

const (
	StatusOK Status = iota
	StatusError
)

type Stringer interface {
	String() string
}

type CreateTodoRequestModel struct {
	Label Stringer
}

type CreateTodoRequestModelResolver interface {
	Resolve() CreateTodoRequestModel
}

type TodoResponseModel struct {
	ID         interface{} `json:"id"`
	Label      string      `json:"label"`
	IsComplete bool        `json:"isComplete"`
}

type TodoResponseModelSender interface {
	Send(TodoResponseModel, Status)
}

type TodosController struct {
	TodoService *todo.Service
}

func (ctrl TodosController) CreateTodo(req CreateTodoRequestModelResolver, res TodoResponseModelSender) {
	model := req.Resolve()

	t, err := ctrl.TodoService.Create(model.Label)
	if err != nil {
		res.Send(TodoResponseModel{}, StatusError)
		return
	}

	res.Send(TodoResponseModel{
		ID:         t.GetID(),
		Label:      t.GetLabel(),
		IsComplete: t.GetIsComplete(),
	}, StatusOK)
}
