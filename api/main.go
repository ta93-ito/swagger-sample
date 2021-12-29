package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", Health)
	http.ListenAndServe(":3000", nil)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	fmt.Fprint(w, "server is healthy!")
}
