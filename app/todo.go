package app

type Stringer interface {
	String() string
}

type String string

func (s String) String() string {
	return string(s)
}

type Status interface {
	Stringer
	Code() uint8
}

type status uint8

func (s status) Code() uint8 {
	return uint8(s)
}

func (s status) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusError:
		return "Error"
	case StatusCreated:
		return "Created"
	}
	return "N/A"
}

const (
	StatusOK status = iota
	StatusError
	StatusCreated
)

type CreateTodoRequestModelResolver interface {
	Resolve() CreateTodoRequestModel
}

type TodoResponseModelSender interface {
	Send(TodoResponseModel, Status)
}

type TodoResponseModel struct {
	id         interface{}
	label      string
	isComplete bool
}

func (m TodoResponseModel) GetID() interface{} {
	return m.id
}

func (m TodoResponseModel) GetLabel() string {
	return m.label
}

func (m TodoResponseModel) GetIsComplete() bool {
	return m.isComplete
}

func NewCreateTodoRequestModel(label Stringer) CreateTodoRequestModel {
	return CreateTodoRequestModel{label.String()}
}

type CreateTodoRequestModel struct {
	label string
}

type TodoService interface {
	Create(label Stringer) (TodoResponseModel, error)
}

func NewTodoController(todoService TodoService) TodoController {
	return TodoController{todoService}
}

type TodoController struct {
	todoService TodoService
}

func (ctrl TodoController) CreateTodo(req CreateTodoRequestModelResolver, res TodoResponseModelSender) {
	model := req.Resolve()

	status := StatusCreated
	todo, err := ctrl.todoService.Create(String(model.label))
	if err != nil {
		status = StatusError
	}

	res.Send(todo, status)
}
