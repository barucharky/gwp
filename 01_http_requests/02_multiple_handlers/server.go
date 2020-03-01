package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}

func world(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "World, Hello!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
