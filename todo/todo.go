package todo

type Status uint8

const (
	StatusOK Status = iota
	StatusError
	StatusCreated
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
}

type Logger interface {
	LogError(error)
}

type TodosController struct {
	Logger
	*TodoService
}

func (ctrl TodosController) CreateTodo(w TodoWriter, label string) {
	todo, err := ctrl.TodoService.CreateTodo(label)
	if err != nil {
		ctrl.Logger.LogError(err)
		w.Write(todo, StatusError)
		return
	}

	w.Write(todo, StatusCreated)
}
