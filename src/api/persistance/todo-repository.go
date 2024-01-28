package persistance

import (
	"errors"
	"fmt"
	"log"

	"github.com/sango-poc-go/src/api/model"
)

var (
	todoMap = make(map[int]model.Todo)
)

func initializeTodos(){
	log.Println("Initializing todo")

	todoMap[1] = model.Todo{
		Id: 1, Detail: "Do taxes", Finished: false,
	}
	todoMap[2] = model.Todo{
		Id: 2, Detail: "Vaccum first floor", Finished: false,
	}
}


func AllTodos() ([]model.Todo, error){
	initializeTodos()

	if len(todoMap) == 0 {
		return nil, errors.New("no todos found")
	}else {
		var todos []model.Todo

		for k := range todoMap {
			todos = append(todos, todoMap[k])
		}
		return todos, nil
	}
}

func GetTodoById(id int ) ([]model.Todo, error){
	initializeTodos()

	if len(todoMap) == 0 {
		return nil, errors.New("no todos found")
	}else {
		var todos []model.Todo
		_, exists := todoMap[id]
		if !exists {
			return nil, fmt.Errorf("no todo with id %d found", id)
		}else {
			todos = append(todos, todoMap[id])
			return todos, nil
		}
	}
}

func AddTodo(newTodo model.Todo) error {
	initializeTodos()

	 _, exists := todoMap[newTodo.Id]
	if exists {
		return fmt.Errorf("todo with id %d exists", newTodo.Id)
	}else {
		todoMap[newTodo.Id] = newTodo
		return nil
	}
}

func UpdateTodo(newTodo model.Todo ) ([]model.Todo, error){
	initializeTodos()

	if len(todoMap) == 0 {
		return nil, errors.New("no todos found")
	}else {
		_, exists := todoMap[newTodo.Id]
		if !exists {
			return nil, fmt.Errorf("no todo with id %d found", newTodo.Id)
		}else {
			todoMap[newTodo.Id] = newTodo

			var todos []model.Todo

			for k ,_ := range todoMap {
				todos = append(todos, todoMap[k])
			}
			return todos, nil
		}
	}
}

func DeleteTodoById(id int) ([]model.Todo, error){
	initializeTodos()

	if len(todoMap) == 0 {
		return nil, errors.New("no todos found")
		}else {
			_, exists := todoMap[id]
			if !exists {
				return nil, fmt.Errorf("no todo with id %d found", id)
			}else {
				delete(todoMap, id)
	
				var todos []model.Todo
	
				for k := range todoMap {
					todos = append(todos, todoMap[k])
				}
				return todos, nil
			}
		}
}