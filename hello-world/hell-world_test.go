package main

import (
	"testing"
)

func TestHelloWorldSuccess(t *testing.T) {
	name := "abc"
	actualResult, _ := HelloWorld(name)

	if actualResult != "Hello World!! "+ name {
		t.Errorf("Did not get the correct result")
	}
}

func TestHelloWorldFailureMessage(t *testing.T) {
	name := ""
	_,  actualError := HelloWorld(name)

	if actualError.Error() != "name can not be empty" {
		t.Errorf("Did not get the correct error")
	}
}