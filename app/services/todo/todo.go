package todo

type Stringer interface {
	String() string
}

type Booler interface {
	Bool() bool
}

type TodosRepository interface {
	InsertTodo(label Stringer, isComplete Booler) (Todo, error)
}

func NewTodo(id interface{}, label string, isComplete bool) Todo {
	return Todo{id, label, isComplete}
}

func newEmptyTodo() Todo {
	return NewTodo(0, "", false)
}

type Todo struct {
	id         interface{}
	label      string
	isComplete bool
}

func (t Todo) GetID() interface{} {
	return t.id
}

func (t Todo) GetLabel() string {
	return t.label
}

func (t Todo) GetIsComplete() bool {
	return t.isComplete
}

type Bool bool

func (b Bool) Bool() bool {
	return bool(b)
}

func NewService(repo TodosRepository) *Service {
	return &Service{repo}
}

type Service struct {
	repo TodosRepository
}

func (srv *Service) Create(label Stringer) (Todo, error) {
	todo, err := srv.repo.InsertTodo(label, Bool(false))
	if err != nil {
		return newEmptyTodo(), err
	}

	return todo, nil
}
