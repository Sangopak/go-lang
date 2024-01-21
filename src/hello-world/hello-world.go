package src

import (
	"errors"
	"fmt"
	"strconv"
)

func HelloWorld(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	}
	response := "Hello World!! "+ name
	fmt.Println(response)
	return response, nil
}

func HelloWorldWithChannel(messages chan string, counter int) {
	fmt.Println("Hello World from other thread start")
	messages <- "This is a message from in the inner thread" + strconv.Itoa(counter)
	fmt.Println("Hello World from other thread end")
	//need to close channel otherwise it creates deadlock as the range operation waits for the channel to close
	close(messages)
}