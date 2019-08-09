package app

type Status = uint8

const (
	StatusOK Status = iota
	StatusError
	StatusCreated
	StatusNotFound
)

type TodoResponseSender interface {
	Send(TodoResponse, Status)
}

type StatusSender interface {
	Send(Status)
}

// TodoResponse is basically a method filter for TodoModel.
// It cannot have any DIFFERENT methods, but can leave some methods out.
type TodoResponse interface {
	GetID() interface{}
	GetLabel() string
	GetIsComplete() bool
}

// TodoModel is what is returned from the service.
type TodoModel interface {
	GetID() interface{}
	GetLabel() string
	SetLabel(string)
	GetIsComplete() bool
	SetIsComplete(bool)
	Save() error
}

type TodoService interface {
	Create(label string) (TodoModel, error)
	GetByID(id interface{}) (TodoModel, error)
	DeleteByID(id interface{}) error
}

type LoggerService interface {
	LogError(error)
}

func NewTodoController(todoService TodoService, loggerService LoggerService) *TodoController {
	return &TodoController{todoService, loggerService}
}

type TodoController struct {
	todoService   TodoService
	loggerService LoggerService
}

type CreateTodoRequest interface {
	Label() string
}

func (ctrl *TodoController) CreateTodo(req CreateTodoRequest, res TodoResponseSender) {
	todo, err := ctrl.todoService.Create(req.Label())
	if err != nil {
		ctrl.loggerService.LogError(err)
		res.Send(nil, StatusError)
		return
	}

	res.Send(todo, StatusCreated)
}

type GetTodoRequest interface {
	ID() interface{}
}

func (ctrl *TodoController) GetTodo(req GetTodoRequest, res TodoResponseSender) {
	todo, err := ctrl.todoService.GetByID(req.ID())
	if err != nil {
		ctrl.loggerService.LogError(err)
		res.Send(nil, StatusError)
		return
	} else if todo == nil {
		res.Send(nil, StatusNotFound)
		return
	}

	res.Send(todo, StatusOK)
}

type DeleteTodoRequest interface {
	ID() interface{}
}

func (ctrl *TodoController) DeleteTodo(req DeleteTodoRequest, res StatusSender) {
	err := ctrl.todoService.DeleteByID(req.ID())
	if err != nil {
		ctrl.loggerService.LogError(err)
		res.Send(StatusError)
		return
	}

	res.Send(StatusOK)
}

type UpdateTodoRequest interface {
	ID() interface{}
	Label() string
	IsComplete() bool
}

func (ctrl *TodoController) UpdateTodo(req UpdateTodoRequest, res TodoResponseSender) {
	todo, err := ctrl.todoService.GetByID(req.ID())
	if err != nil {
		ctrl.loggerService.LogError(err)
		res.Send(nil, StatusError)
		return
	} else if todo == nil {
		res.Send(nil, StatusNotFound)
		return
	}

	todo.SetLabel(req.Label())
	todo.SetIsComplete(req.IsComplete())

	if err = todo.Save(); err != nil {
		ctrl.loggerService.LogError(err)
		res.Send(nil, StatusError)
		return
	}

	res.Send(todo, StatusOK)
}
