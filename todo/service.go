package todo

type TodoService struct {
	TodoRepository
}

func (s TodoService) CreateTodo(label string) (Todo, error) {
	todo := Todo{Label: label, IsComplete: false}

	id, err := s.TodoRepository.Insert(todo.Label, todo.IsComplete)
	if err != nil {
		return todo, err
	}

	// Set the id
	todo.ID = id

	return todo, nil
}
