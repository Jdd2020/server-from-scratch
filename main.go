package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const PORT = ":8080"
const SERVER_ADDR_ROOT = "http://localhost"
const FILE_PATH_ROOT = "."
const IMAGE_PATH = "/assets"
const URL_PREFIX = "/app"

type apiConfig struct {
	fileServerHits atomic.Int32
}


func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) resetMetrics(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	cfg.fileServerHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (cfg *apiConfig) serveMetrics(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	resp := fmt.Sprintf("Hits: %d", cfg.fileServerHits.Load())
	log.Printf("Serving metrics with %d hits", cfg.fileServerHits.Load())
	w.Write([]byte(resp))
}

func main() {
	const URL_ROOT = SERVER_ADDR_ROOT + PORT + URL_PREFIX

	mux := http.NewServeMux()

	cfg := apiConfig{fileServerHits : atomic.Int32{}}


	dir := http.Dir(FILE_PATH_ROOT)
	mux.Handle("/app/", cfg.middlewareMetricsInc((http.StripPrefix("/app", http.FileServer(dir)))))
	mux.HandleFunc("/healthz", healthzHandler)
	mux.HandleFunc("/metrics", cfg.serveMetrics)
	mux.HandleFunc("/reset", cfg.resetMetrics)
	

	srv := http.Server{
		Addr: PORT,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler: mux,
	}

	log.Printf("Server running on %s\n", URL_ROOT)
	log.Fatal(srv.ListenAndServe())
}

func healthzHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}