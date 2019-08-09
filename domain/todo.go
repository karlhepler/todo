package domain

type TodoModel interface {
	GetID() interface{}
	GetLabel() string
	SetLabel(string)
	GetIsComplete() bool
	SetIsComplete(bool)
	Save() error
}

type TodoRepository interface {
	DeleteByID(interface{}) error
	GetByID(interface{}) (TodoModel, error)
	Create(label string, isComplete bool) (TodoModel, error)
}

func NewTodoService(todoRepository TodoRepository) *TodoService {
	return &TodoService{todoRepository}
}

type TodoService struct {
	todoRepository TodoRepository
}

func (srv *TodoService) Create(label string) (TodoModel, error) {
	return srv.todoRepository.Create(label, false)
}

func (srv *TodoService) GetByID(id interface{}) (TodoModel, error) {
	return srv.todoRepository.GetByID(id)
}

func (srv *TodoService) DeleteByID(id interface{}) error {
	return srv.todoRepository.DeleteByID(id)
}
