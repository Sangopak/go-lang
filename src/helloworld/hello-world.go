package src

import (
	"fmt"
	"errors"
)

func HelloWorld(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	}
	response := "Hello World!! "+ name
	fmt.Println(response)
	return response, nil
}