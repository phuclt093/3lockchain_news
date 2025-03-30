package handlers

import (
	"fmt"
	"net/http"
)

// Handler cho route /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ChainNews API Gateway is running!")
}
