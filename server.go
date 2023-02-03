package main

import (
	"fmt"
	"net/http"
)

type msg string

var counter int

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
	counter = counter + 1
	fmt.Fprint(resp, counter)

}

func main() {
	msgHandler := msg("Hello from Web Server in Go 8")
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", msgHandler)
}
