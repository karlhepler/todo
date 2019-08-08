package driver

type TodosRepository interface {
	Create(label string, isComplete bool) (interface{}, error)
	UpdateByID(id interface{}, label string, isComplete bool) error
	DeleteByID(id interface{}) error
	GetByID(id interface{}, label *string, isComplete *bool) error
	ForEach(func(id interface{}, label string, isComplete bool)) error
}

type Logger interface {
	Printf(format string, v ...interface{})
}
