package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "{\"message\":\"hello\"}")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    const port string = "8090"
    const host string = ":"

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    fmt.Println("server starting at", port)
    http.ListenAndServe(host + port, nil)
}