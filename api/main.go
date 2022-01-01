package main

import (
	"net/http"

	"github.com/d4l3k/go-pry/pry"
	"github.com/ta93-ito/golang-swagger-sample/api/handler"
)

func main() {
	pry.Pry()
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/todo", handler.POSTTODO)
	mux.HandleFunc("/todos", handler.GETAllTODOs)
	http.ListenAndServe(":3000", mux)
}
