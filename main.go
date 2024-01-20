package main

import (
	helloworld "github.com/sango-poc-go/src/helloworld"
	threadExample "github.com/sango-poc-go/src/threads"
)

func main(){
	helloworld.HelloWorld("Sango")
	threadExample.ThreadExample()
}