package main

import (
	"fmt"
	"strconv"
	//"time"
)

func main(){
	fmt.Println("This is from main thread start")
	messages := make(chan string)
	var counter int;
	for counter =0; counter <10; counter++ {
		go HelloWorld(messages, counter)
		fmt.Println("This is from main with counter ", counter)
	}
	fmt.Println("This is from main thread end")
	//time.Sleep(2*time.Second) //mimicing some other important stuff to happen

	// for range loop to go over all the messages in the channel
	for msg := range messages {
		fmt.Println("Message from the thread",msg)
	}
    
}

func HelloWorld(messages chan string, counter int) {
	fmt.Println("Hello World from other thread start")
	messages <- "This is a message from in the inner thread" + strconv.Itoa(counter)
	fmt.Println("Hello World from other thread end")
	//need to close channel otherwise it creates deadlock as the range operation waits for the channel to close
	close(messages)
}