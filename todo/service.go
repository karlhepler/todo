package todo

type TodoFactory struct {
	//
}

func (f TodoFactory) Make(label string, isComplete bool) Todo {
	return Todo{
		Label:      label,
		IsComplete: isComplete,
	}
}

type TodoService struct {
	TodoRepository
	*TodoFactory
}

func (s TodoService) CreateTodo(label string) (Todo, error) {
	todo := s.TodoFactory.Make(label, false)

	id, err := s.TodoRepository.Insert(todo.Label, todo.IsComplete)
	if err != nil {
		return todo, err
	}

	// Set the id
	todo.ID = id

	return todo, nil
}
