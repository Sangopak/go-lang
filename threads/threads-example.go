package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("This is from main thread start")
	messages := make(chan string)
	go HelloWorld(messages)
	fmt.Println("This is from main thread end")
	time.Sleep(2*time.Second) //Give sometime to run the other threads
	msg := <-messages
    fmt.Printf("Message from the thread %s",msg)
}

func HelloWorld(messages chan string) {
	fmt.Println("Hello World from other thread start")
	messages <- "This is a message from in the inner thread"
	fmt.Println("Hello World from other thread end")
}