package domain

type Stringer interface {
	String() string
}

type Booler interface {
	Bool() bool
}

type Todo interface {
	GetID() interface{}
	GetLabel() string
	GetIsComplete() bool
}

type TodoModel struct {
	id         interface{}
	label      string
	isComplete bool
}

func (m TodoModel) GetID() interface{} {
	return m.id
}

func (m TodoModel) GetLabel() string {
	return m.label
}

func (m TodoModel) GetIsComplete() bool {
	return m.isComplete
}

type TodosRepository interface {
	Save(Todo) error
	DeleteByID(interface{}) error
	GetByID(interface{}) (TodoModel, error)
}

func NewTodoService(todosRepo TodosRepository) TodoService {
	return TodoService{todosRepo}
}

type TodoService struct {
	todosRepo TodosRepository
}

func (s TodoService) Create(label Stringer) (TodoModel, error) {
	todoModel := TodoModel{label: label.String()}

	if err := s.todosRepo.Save(todoModel); err != nil {
		return TodoModel{}, err
	}

	return todoModel, nil
}

func (s TodoService) Update(todo Todo) (TodoModel, error) {
	if err := s.todosRepo.Save(todo); err != nil {
		return TodoModel{}, err
	}

	return TodoModel{
		id:         todo.GetID(),
		label:      todo.GetLabel(),
		isComplete: todo.GetIsComplete(),
	}, nil
}

func (s TodoService) GetByID(id interface{}) (TodoModel, error) {
	return s.todosRepo.GetByID(id)
}

func (s TodoService) DeleteByID(id interface{}) error {
	return s.todosRepo.DeleteByID(id)
}
