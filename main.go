package main

import (
	"log"
	"net/http"
	"time"
)

const PORT = ":8080"
const SERVER_ADDR_ROOT = "http://localhost"
const FILE_PATH_ROOT = "."
const IMAGE_PATH = "/assets"
const URL_PREFIX = "/app"



func main() {
	const URL_ROOT = SERVER_ADDR_ROOT + PORT + URL_PREFIX

	mux := http.NewServeMux()


	dir := http.Dir(FILE_PATH_ROOT)
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(dir) ))
	mux.HandleFunc("/healthz", healtzHandler)
	

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

func healtzHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}