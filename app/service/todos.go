package service

import (
	"github.com/oldtimeguitarguy/todo/app/command/service"
	"github.com/oldtimeguitarguy/todo/app/service/driver"
)

type Todos struct {
	driver.TodosRepository
}

func (srv Todos) Create(label string) (service.TodoModel, error) {
	const isComplete = false

	id, err := srv.TodosRepository.Create(label, isComplete)
	if err != nil {
		return todoModel{}, err
	}

	return todoModel{
		id:         id,
		label:      label,
		isComplete: isComplete,
	}, nil
}

func (srv Todos) UpdateByID(id interface{}, label string, isComplete bool) (service.TodoModel, error) {
	err := srv.TodosRepository.UpdateByID(id, label, isComplete)
	if err != nil {
		return todoModel{}, err
	}

	return todoModel{
		id:         id,
		label:      label,
		isComplete: isComplete,
	}, nil
}

func (srv Todos) DeleteByID(id interface{}) error {
	return srv.TodosRepository.DeleteByID(id)
}

func (srv Todos) GetByID(id interface{}) (service.TodoModel, error) {
	todo := todoModel{id: id}

	err := srv.TodosRepository.GetByID(id, &todo.label, &todo.isComplete)
	if err != nil {
		return todoModel{}, err
	}

	return todo, nil
}

func (srv Todos) GetList() (service.TodoListModel, error) {
	todoList := todoListModel{}

	err := srv.TodosRepository.ForEach(func(id interface{}, label string, isComplete bool) {
		todoList = append(todoList, todoModel{
			id:         id,
			label:      label,
			isComplete: isComplete,
		})
	})
	if err != nil {
		return todoListModel{}, err
	}

	return todoList, nil
}

type todoModel struct {
	id         interface{}
	label      string
	isComplete bool
}

func (mdl todoModel) GetID() interface{} {
	return mdl.id
}

func (mdl todoModel) GetLabel() string {
	return mdl.label
}

func (mdl todoModel) GetIsComplete() bool {
	return mdl.isComplete
}

type todoListModel []todoModel

func (mdl todoListModel) GetLength() int {
	return len(mdl)
}

func (mdl todoListModel) ForEach(callback func(todo service.TodoModel, index int)) {
	for index, todo := range mdl {
		callback(todo, index)
	}
}
