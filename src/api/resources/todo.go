package resources

import (
	"log"
	"net/http"
	"strconv"

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