package main

import (
	"fmt"
	"net/http"
)

func pinky(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "NARF!")
}

func brain(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Try to take over the world!")
}

func main() {

	// -- ------------------------------
	/*
			Defines parameters for running an HTTP server.
			The Server struct configuration:
		    type Server struct {
		        Addr           string
		        Handler        Handler
		        ReadTimeout    time.Duration
		        WriteTimeout   time.Duration
		        MaxHeaderBytes int
		        TLSConfig      *tls.Config
		        TLSNextProto   map[string]func(*Server, *tls.Conn, Handler)
		        ConnState      func(net.Conn, ConnState)
		        ErrorLog       *log.Logger
		    }
			For details see: https://pkg.go.dev/net/http?tab=doc#Server
	*/

	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.HandleFunc("/pinky", pinky)
	http.HandleFunc("/brain", brain)

	server.ListenAndServe()
}
