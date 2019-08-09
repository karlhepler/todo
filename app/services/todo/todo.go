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

type BoolBooler bool

func (b BoolBooler) Bool() bool {
	return bool(b)
}

type Service struct {
	Repo TodosRepository
}

func (srv Service) Create(label Stringer) (Todo, error) {
	const isComplete = false

	todo, err := srv.Repo.InsertTodo(label, BoolBooler(isComplete))
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}
