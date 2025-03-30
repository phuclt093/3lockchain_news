package main

import (
	"chainnews_api_gateway/handlers" // Import package handlers
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Định nghĩa các route và gọi handler từ package handlers
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")

	port := "8000"
	fmt.Println("API Gateway is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
