package thread

import (
	"fmt"
	helloworld "github.com/sango-poc-go/src/hello-world"
	//"time"
)

func ThreadExample(){
	fmt.Println("This is from main thread start")
	messages := make(chan string)
	var counter int;
	for counter =0; counter <10; counter++ {
		go helloworld.HelloWorldWithChannel(messages, counter)
		fmt.Println("This is from main with counter ", counter)
	}
	fmt.Println("This is from main thread end")
	//time.Sleep(2*time.Second) //mimicing some other important stuff to happen

	// for range loop to go over all the messages in the channel
	for msg := range messages {
		fmt.Println("Message from the thread",msg)
	}
    
}
