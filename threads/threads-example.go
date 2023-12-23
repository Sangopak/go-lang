package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("This is from main thread start")
	go HelloWorld()
	fmt.Println("This is from main thread end")
	time.Sleep(2*time.Second) //Give sometime to run the other threads
}

func HelloWorld() {
	fmt.Println("Hello World from other thread start")
	time.Sleep(time.Second)
	fmt.Println("Hello World from other thread end")
}