package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!, service started")
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe(":8080", nil)
}
