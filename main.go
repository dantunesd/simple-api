package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", helloWordHandler)
	http.ListenAndServe(":8080", serveMux)
}

func helloWordHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %s Path: %s", r.Method, r.URL.Path)

	type Body struct {
		Message string `json:"message"`
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Body{"hello world"})
}
