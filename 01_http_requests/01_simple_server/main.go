package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(
		writer,
		"Hello, %s!",
		request.URL.Path[1:],
	)
}

func main() {
	http.HandleFunc("/", handler)

	// -- ------------------------------
	/*
		func ListenAndServe(addr string, handler Handler) error
			- It listens on the TCP network `addr`
			- It then calls Serve with `handler` to handle requests on incoming connections.
		The handler is typically nil, in which case the DefaultServeMux is used.
	*/

	http.ListenAndServe(":8090", nil)
}
