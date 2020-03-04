package main

import (
	"fmt"
	"net/http"
)

func narf(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "NARF!")
}

func pondering(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Are you pondering what I'm pondering?")
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Instead of %s, try /pinky or /brain", request.URL.Path[1:])
}

func main() {

	var server http.Server = http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/pinky", narf)
	http.HandleFunc("/brain", pondering)
	http.HandleFunc("/", handler)

	server.ListenAndServe()
}
