package response

type Status uint8

const (
	StatusOK Status = iota
	StatusError
	StatusCreated
)

type Sender interface {
	Send(Status)
}

type TodoSender interface {
	Send(Todo, Status)
}

type TodoListSender interface {
	Send(TodoList, Status)
}

type Todo struct {
	ID         interface{}
	Label      string
	IsComplete bool
}

type TodoList []Todo
