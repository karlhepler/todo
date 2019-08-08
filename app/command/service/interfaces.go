package service

type Todos interface {
	Create(label string) (TodoModel, error)
	UpdateByID(id interface{}, label string, isComplete bool) (TodoModel, error)
	DeleteByID(id interface{}) error
	GetByID(id interface{}) (TodoModel, error)
	GetList() (TodoListModel, error)
}

type Logger interface {
	LogError(error)
}

type TodoModel interface {
	GetID() interface{}
	GetLabel() string
	GetIsComplete() bool
}

type TodoListModel interface {
	GetLength() int
	ForEach(func(todo TodoModel, index int))
}
