package main

import (
	"fmt"
	"net/http"

	"github.com/ta93-ito/golang-swagger-sample/api/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/todo", handler.POSTTODO)
	mux.HandleFunc("/todos", handler.GETAllTODOs)
	http.ListenAndServe(":3000", mux)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	fmt.Fprint(w, "server is healthy!")
}
