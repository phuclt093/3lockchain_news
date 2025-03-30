package handlers

import (
	"fmt"
	"net/http"
)

// Handler cho route /hello
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message": "Hello from ChainNews API Gateway!"}`)
}
