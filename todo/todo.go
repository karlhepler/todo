package todo

type Status uint8

const (
	StatusOK Status = iota
	StatusError
	StatusCreated
	StatusNotFound
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
	GetByID(id interface{}, label *string, isComplete *bool) error
}

type Logger interface {
	LogError(error)
}

func NewTodosController(logger Logger, todoRepository TodoRepository) *TodosController {
	return &TodosController{
		logger,
		todoService{todoRepository},
	}
}

type TodosController struct {
	logger      Logger
	todoService todoService
}

func (ctrl TodosController) CreateTodo(w TodoWriter, label string) {
	todo, err := ctrl.todoService.createTodo(label)
	if err != nil {
		ctrl.logger.LogError(err)
		w.Write(todo, StatusError)
		return
	}

	w.Write(todo, StatusCreated)
}

func (ctrl TodosController) GetTodoByID(w TodoWriter, id interface{}) {
	todo, err := ctrl.todoService.getTodoByID(id)
	if err != nil {
		ctrl.logger.LogError(err)
		w.Write(todo, StatusError)
		return
	}

	w.Write(todo, StatusOK)
}
