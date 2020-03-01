package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(
		writer,
		"Hello world, %s!",
		request.URL.Path[1:],
	)
}

func pinky(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(
		writer,
		"NARF!",
	)
}

func brain(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(
		writer,
		"Same thing we do every night, Pinky; try to take over the world!",
	)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/pinky", pinky)
	http.HandleFunc("/brain", brain)
	http.ListenAndServe(":8090", nil)
}
