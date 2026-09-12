package main

import (
	"net/http"
	"time"
)

const PORT = ":8080"

func main() {
	mux := http.NewServeMux()

	srv := http.Server{
		Addr: PORT,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler: http.Handler(mux),
	}

	srv.ListenAndServe()
}