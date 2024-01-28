package main

import (
	//helloworld "github.com/sango-poc-go/src/helloworld"
	//threadExample "github.com/sango-poc-go/src/threads"
	// httpServer "github.com/sango-poc-go/src/http/server"
	// httpClient "github.com/sango-poc-go/src/http/client"

	resources "github.com/sango-poc-go/src/api/resources"
	ginServer "github.com/sango-poc-go/src/api/server"
)

func main() {
	//simple hello world
	//helloworld.HelloWorld("Sango")

	//simple thread with counter
	//threadExample.ThreadExample()

	// run the server in a different thread once the client code runs it will be stopped
	// go httpServer.HttpServerExample()
	// httpClient.HttpClientExample()

	//run a gin server
	router := ginServer.CreateDefaultRouter()
	router.GET("/", resources.GetHeatlhCheck)
	router.GET("/todos", resources.GetTodos)
	router.GET("/todos/:id", resources.GetTodosById)

	router.Run()
}
