package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/oldtimeguitarguy/todo/app/command"
	"github.com/oldtimeguitarguy/todo/app/command/response"
	"github.com/oldtimeguitarguy/todo/app/service"
)

type Services struct {
	Todos  func() service.Todos
	Logger func() service.Logger
}

type Application struct {
	CreateTodo  func() command.CreateTodo
	GetTodoList func() command.GetTodoList
}

var srv = &Services{}
var app = &Application{}

func init() {
	srv.Logger = func() service.Logger {
		return service.Logger{
			Logger: log.New(os.Stderr, "", log.LstdFlags),
		}
	}
	srv.Todos = func() service.Todos {
		return service.Todos{
			TodosRepository: &TodoMemoryRepository{},
		}
	}

	app.CreateTodo = func() command.CreateTodo {
		return command.CreateTodo{
			Todos:  srv.Todos(),
			Logger: srv.Logger(),
		}
	}
	app.GetTodoList = func() command.GetTodoList {
		return command.GetTodoList{
			Todos:  srv.Todos(),
			Logger: srv.Logger(),
		}
	}
}

func main() {
	router := httprouter.New()

	router.POST("/todos", storeTodoHandler)
	router.GET("/todos", listTodosHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func storeTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := command.CreateTodoRequest{r.FormValue("label")}
	res := JSONTodoSender{w}

	app.CreateTodo().Call(req, res)
}

func listTodosHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := command.GetTodoListRequest{}
	res := JSONTodoListSender{w}

	app.GetTodoList().Call(req, res)
}

type JSONTodoSender struct {
	http.ResponseWriter
}

func (res JSONTodoSender) Send(todo response.Todo, status response.Status) {
	if status != response.StatusCreated {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	data, err := json.Marshal(map[string]interface{}{
		"id":         todo.ID,
		"label":      todo.Label,
		"isComplete": todo.IsComplete,
	})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write(data)
}

type JSONTodoListSender struct {
	http.ResponseWriter
}

func (res JSONTodoListSender) Send(todoList response.TodoList, status response.Status) {
	if status != response.StatusOK {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	data, err := json.Marshal(todoList)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

type TodoMemoryRepository struct {
	todos []map[string]string
}

func (repo *TodoMemoryRepository) Create(label string, isComplete bool) (interface{}, error) {
	var err error
	var lastID = 0
	if len(repo.todos) > 0 {
		lastTodo := repo.todos[len(repo.todos)-1]
		if lastID, err = strconv.Atoi(lastTodo["id"]); err != nil {
			return nil, err
		}
	}

	todo := map[string]string{
		"id":         strconv.Itoa(lastID + 1),
		"label":      label,
		"isComplete": strconv.FormatBool(isComplete),
	}

	repo.todos = append(repo.todos, todo)

	return todo["id"], nil
}

func (repo *TodoMemoryRepository) UpdateByID(id interface{}, label string, isComplete bool) error {
	if len(repo.todos) == 0 {
		return fmt.Errorf("Unable to find todo with id %s", id)
	}

	lastTodo := repo.todos[len(repo.todos)-1]
	lastID, err := strconv.Atoi(lastTodo["id"])
	if err != nil {
		return fmt.Errorf("Something went wrong converting id to int")
	}

	var index int
	var todo map[string]string
	isFound := false
	for index, todo = range repo.todos {
		if todo["id"] == id {
			isFound = true
			break
		}
	}

	if isFound == false {
		return fmt.Errorf("Unable to find todo with id %s", id)
	}

	repo.todos = append(repo.todos[:index], repo.todos[index+1:]...)
	repo.todos = append(repo.todos, map[string]string{
		"id":         strconv.Itoa(lastID + 1),
		"label":      label,
		"isComplete": strconv.FormatBool(isComplete),
	})

	return nil
}

func (repo *TodoMemoryRepository) DeleteByID(id interface{}) error {
	if len(repo.todos) == 0 {
		return fmt.Errorf("Unable to find todo with id %s", id)
	}

	var index int
	var todo map[string]string
	isFound := false
	for index, todo = range repo.todos {
		if todo["id"] == id {
			isFound = true
			break
		}
	}

	if isFound == false {
		return fmt.Errorf("Unable to find todo with id %s", id)
	}

	repo.todos = append(repo.todos[:index], repo.todos[index+1:]...)

	return nil
}

func (repo *TodoMemoryRepository) GetByID(id interface{}, label *string, isComplete *bool) error {
	if len(repo.todos) == 0 {
		return fmt.Errorf("Unable to find todo with id %s", id)
	}

	var todo map[string]string
	isFound := false
	for _, todo = range repo.todos {
		if todo["id"] == id {
			isFound = true
			break
		}
	}

	if isFound == false {
		return fmt.Errorf("Unable to find todo with id %s", id)
	}

	var err error
	*label = todo["label"]
	*isComplete, err = strconv.ParseBool(todo["isComplete"])

	return err
}

func (repo *TodoMemoryRepository) ForEach(callback func(id interface{}, label string, isComplete bool)) error {
	if len(repo.todos) == 0 {
		return nil
	}

	for _, todo := range repo.todos {
		isComplete, err := strconv.ParseBool(todo["isComplete"])
		if err != nil {
			return err
		}
		callback(todo["id"], todo["label"], isComplete)
	}

	return nil
}
