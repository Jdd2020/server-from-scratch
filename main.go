package main

import (
	"log"
	"net/http"
	"time"
)

const PORT = ":8080"
const SERVER_ADDR_ROOT = "http://localhost"
const FILE_PATH_ROOT = "."

func main() {
	mux := http.NewServeMux()

	/* Serve HTML files from the root directory */
	directory := http.Dir(FILE_PATH_ROOT)
	fileServer := http.FileServer(directory)
	mux.Handle("/", fileServer)
	

	srv := http.Server{
		Addr: PORT,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler: mux,
	}

	log.Printf("Server running on %s\n", SERVER_ADDR_ROOT + PORT)
	srv.ListenAndServe()
}