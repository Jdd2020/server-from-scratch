package main

import (
	"net/http"
	"time"
)

const PORT = ":8080"

func main() {
	mux := http.NewServeMux()
	
	/* Set file directory to root */ 
	directory := http.Dir("./")

	/* Create file handler/server with root directory */ 
	fileServer := http.FileServer(directory)

	/* Set base path to serve static html using the file server */
	mux.Handle("/", fileServer)
	

	srv := http.Server{
		Addr: PORT,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler: mux,
	}

	srv.ListenAndServe()
}