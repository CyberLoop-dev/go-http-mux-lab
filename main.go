package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

type response struct {
	Route string `json:"route"`
	Method string `json:"method"`
	Message string `json:"message"`
	Values map[string]string `json:"values,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, value response) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, response{Route: "root", Method: r.Method, Message: "exact root route"})
	})
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, response{Route: "health", Method: r.Method, Message: "ok"})
	})
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, response{Route: "user", Method: r.Method, Message: "matched one path segment", Values: map[string]string{"id": r.PathValue("id")}})
	})
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusCreated, response{Route: "create-user", Method: r.Method, Message: "POST matched the collection route"})
	})
	mux.HandleFunc("GET /files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, response{Route: "file-path", Method: r.Method, Message: "matched trailing wildcard", Values: map[string]string{"path": r.PathValue("path")}})
	})
	return mux
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "local listen address")
	flag.Parse()
	log.Printf("Go ServeMux lab listening on http://%s", *addr)
	log.Fatal(http.ListenAndServe(*addr, newMux()))
}
