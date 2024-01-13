package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
)

func main() {

    resp, err := GetHttp("http://localhost:8090/hello")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func GetHttp(endpoint string) (*http.Response, error){
    if endpoint == "" {
        return nil, errors.New("endpoint can not be empty")
    }
    return http.Get(endpoint)
}