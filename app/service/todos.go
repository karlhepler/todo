package service

import (
	"github.com/oldtimeguitarguy/todo/app/command/service"
	"github.com/oldtimeguitarguy/todo/app/service/driver"
)

type Todos struct {
	driver.TodosRepository
}

// The functions MUST take in interfaces and spit out non-interfaces.
// The non-interfaces will contain data and services and stuff that were attached to them by the function.
// These could be private things if they need to be.
// The objects can have functions attached to them, which is just a way of moving functions around with data and state.
// If you expect the data to go through all of the functions at some point, then they're all worth having.
// So.........
// In this case, I could have done something like...

// NO SERVICE INTERFACES.
// THE COMMANDS AND SERVICES ARE TIGHTLY COUPLED ON PURPOSE.
// THEY ACT AS ONE THING. THEY ARE CO-DEPENDENT.

// On the other hand, services work as an airlock.
// They must be built with structs containing ONLY interfaces.

// Do they need to be structs at all?
// Can't they all just be independent modules with functions?
// You can't move the module around.
// That's ok - you don't need to. Just import it into the file that needs it.
// Change the import string or the code behind it to change the implementation.

// WAIT! Switching the import path CANNOT be done at runtime.
// This is where interfaces come in.

// The rule "functions take in interfaces and spit out non-interfaces"
// forces you to use interfaces... which means you might end up using a struct.
// BUT you don't have to. I used a string, for instance.
// Since the interface only had one method, it was straight-forward.

type Labeler interface {
	Label() string
}

func (srv Todos) Create(l Labler) (TodoModel, error) {
	label := l.Label()
}

////

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
