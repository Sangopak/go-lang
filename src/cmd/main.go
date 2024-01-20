package main

import (
	//helloworld "github.com/sango-poc-go/src/helloworld"
	//threadExample "github.com/sango-poc-go/src/threads"
	httpServer "github.com/sango-poc-go/src/http/server"
	httpClient "github.com/sango-poc-go/src/http/client"
)

func main(){
	//simple hello world
	//helloworld.HelloWorld("Sango")
	//simple thread with counter
	//threadExample.ThreadExample()
	// run the server in a different thread once the client code runs it will be stopped
	go httpServer.HttpServerExample() 
	httpClient.HttpClientExample()
}