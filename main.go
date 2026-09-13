package main

import (
	"log"
	"net/http"
	"time"
)

const PORT = ":8080"
const SERVER_ADDR_ROOT = "http://localhost"
const FILE_PATH_ROOT = "."
const IMAGE_PATH = "/assets/logo.png"

func main() {
	const URL_ROOT = SERVER_ADDR_ROOT + PORT

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

	log.Printf("Server running on %s\n", URL_ROOT)
	log.Printf("Fun image availabile at %s\n", URL_ROOT + IMAGE_PATH)
	log.Fatal(srv.ListenAndServe())
}