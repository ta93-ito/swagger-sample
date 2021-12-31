package main

import (
	"net/http"

	"github.com/ta93-ito/golang-swagger-sample/api/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/todo", handler.POSTTODO)
	mux.HandleFunc("/todos", handler.GETAllTODOs)
	http.ListenAndServe(":3000", mux)
}
