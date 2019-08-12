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
	TodoRepository
}

func (ctrl TodosController) CreateTodo(w TodoWriter, label string) {
	// Make the todo
	todo := Todo{
		Label:      label,
		IsComplete: false,
	}

	id, err := ctrl.Insert(todo.Label, todo.IsComplete)
	if err != nil {
		ctrl.Logger.LogError(err)
		w.Write(todo, StatusError)
		return
	}

	// Set the id
	todo.ID = id

	w.Write(todo, StatusCreated)
}
