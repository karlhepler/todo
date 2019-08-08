package command

import (
	"github.com/oldtimeguitarguy/todo/app/command/response"
	"github.com/oldtimeguitarguy/todo/app/command/service"
)

type CreateTodo struct {
	service.Todos
	service.Logger
}

type CreateTodoRequest struct {
	Label string
}

func (cmd CreateTodo) Call(req CreateTodoRequest, res response.TodoSender) {
	todo, err := cmd.Todos.Create(req.Label)
	if err != nil {
		cmd.Logger.LogError(err)
		res.Send(response.Todo{}, response.StatusError)
		return
	}

	res.Send(response.Todo{
		ID:         todo.GetID(),
		Label:      todo.GetLabel(),
		IsComplete: todo.GetIsComplete(),
	}, response.StatusCreated)
}

type UpdateTodo struct {
	service.Todos
	service.Logger
}

type UpdateTodoRequest struct {
	ID         interface{}
	Label      string
	IsComplete bool
}

func (cmd UpdateTodo) Call(req UpdateTodoRequest, res response.TodoSender) {
	todo, err := cmd.Todos.UpdateByID(req.ID, req.Label, req.IsComplete)
	if err != nil {
		cmd.Logger.LogError(err)
		res.Send(response.Todo{}, response.StatusError)
		return
	}

	res.Send(response.Todo{
		ID:         todo.GetID(),
		Label:      todo.GetLabel(),
		IsComplete: todo.GetIsComplete(),
	}, response.StatusOK)
}

type DeleteTodo struct {
	service.Todos
	service.Logger
}

type DeleteTodoRequest struct {
	ID interface{}
}

func (cmd DeleteTodo) Call(req DeleteTodoRequest, res response.Sender) {
	if err := cmd.Todos.DeleteByID(req.ID); err != nil {
		cmd.Logger.LogError(err)
		res.Send(response.StatusError)
		return
	}

	res.Send(response.StatusOK)
}

type ShowTodo struct {
	service.Todos
	service.Logger
}

type ShowTodoRequest struct {
	ID interface{}
}

func (cmd ShowTodo) Call(req ShowTodoRequest, res response.TodoSender) {
	todo, err := cmd.Todos.GetByID(req.ID)
	if err != nil {
		cmd.Logger.LogError(err)
		res.Send(response.Todo{}, response.StatusError)
		return
	}

	res.Send(response.Todo{
		ID:         todo.GetID(),
		Label:      todo.GetLabel(),
		IsComplete: todo.GetIsComplete(),
	}, response.StatusOK)
}

type ListTodos struct {
	service.Todos
	service.Logger
}

type ListTodosRequest struct {
	//
}

func (cmd ListTodos) Call(req ListTodosRequest, res response.TodoListSender) {
	todoList, err := cmd.Todos.GetList()
	if err != nil {
		cmd.Logger.LogError(err)
		res.Send(response.TodoList{}, response.StatusError)
		return
	}

	todoListResponse := make(response.TodoList, todoList.GetLength())
	todoList.ForEach(func(todo service.TodoModel, index int) {
		todoListResponse[index] = response.Todo{
			ID:         todo.GetID(),
			Label:      todo.GetLabel(),
			IsComplete: todo.GetIsComplete(),
		}
	})

	res.Send(todoListResponse, response.StatusOK)
}
