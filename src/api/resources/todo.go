package resources

import (
	"log"
	"net/http"
	"strconv"

	"github.com/sango-poc-go/src/api/model"
	"github.com/sango-poc-go/src/api/persistance"

	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {
	todos, err := persistance.AllTodos()
	if err != nil {
		log.Printf("Something went wrong %s", err)
		context.IndentedJSON(http.StatusNotFound, gin.H{"errorMessage": string(err.Error())})
	} else {
		log.Println("Data found")
		context.IndentedJSON(http.StatusOK, todos)
	}
}

func GetTodosById(context *gin.Context) {
	id,errClient := strconv.Atoi(context.Param("id"))
	if errClient != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": string(errClient.Error())})	
	}

	todos, err := persistance.GetTodoById(id)
	if err != nil {
		log.Printf("Something went wrong %s", err)
		context.IndentedJSON(http.StatusNotFound, gin.H{"errorMessage": string(err.Error())})
	} else {
		log.Println("Data found")
		context.IndentedJSON(http.StatusOK, todos)
	}
}

func PostTodo(context *gin.Context){
	var newTodo model.Todo

	if errClient := context.BindJSON(&newTodo); errClient != nil {
		log.Printf("Something went wrong %s", errClient)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": string(errClient.Error())})
	}else {
		err := persistance.AddTodo(newTodo)

		if err != nil {
			log.Printf("Something went wrong %s", err)
			context.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": string(err.Error())})
		}else {
			log.Println("Data added")
			context.IndentedJSON(http.StatusCreated, newTodo)
		}
	}

}

func UpdateTodo(context *gin.Context){
	var newTodo model.Todo

	if errClient := context.BindJSON(&newTodo); errClient != nil {
		log.Printf("Something went wrong %s", errClient)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": string(errClient.Error())})
	}else {
		todos, err := persistance.UpdateTodo(newTodo)

		if err != nil {
			log.Printf("Something went wrong %s", err)
			context.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": string(err.Error())})
		}else {
			log.Println("Data added")
			context.IndentedJSON(http.StatusOK, todos)
		}
	}

}

func DeleteTodoById(context *gin.Context) {
	id,errClient := strconv.Atoi(context.Param("id"))
	if errClient != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"errorMessage": string(errClient.Error())})	
	}

	todos, err := persistance.DeleteTodoById(id)
	if err != nil {
		log.Printf("Something went wrong %s", err)
		context.IndentedJSON(http.StatusNotFound, gin.H{"errorMessage": string(err.Error())})
	} else {
		log.Println("Data found")
		context.IndentedJSON(http.StatusOK, todos)
	}
}