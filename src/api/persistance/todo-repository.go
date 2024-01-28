package persistance

import (
	"errors"
	"fmt"

	"github.com/sango-poc-go/src/api/model"
)

var todos = []model.Todo{
	{Id: 1, Detail: "Do taxes", Finished: false},
	{Id: 2, Detail: "Vaccum first floor", Finished: false},
}

func AllTodos() ([]model.Todo, error){
	if len(todos) == 0 {
		return nil, errors.New("no todos found")
	}else {
		return todos, nil
	}
}

func GetTodoById(id int ) ([]model.Todo, error){
	if len(todos) == 0 {
		return nil, errors.New("no todos found")
	}else {
		todo := []model.Todo{}
		for _, t := range todos {
			if t.Id == id {
				todo = append(todo, t)
			}
		}
		if len(todo) == 0 {
			return nil, fmt.Errorf("no todo with id %d found", id)
		} else {
			return todo, nil
		}
	}
}

func AddTodo(newTodo model.Todo) error {
	for _, t := range todos {
		if t.Id == newTodo.Id {
			return fmt.Errorf("todo with id %d exists", newTodo.Id)
		}
	}
	todos = append(todos, newTodo)
	return nil
}