package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	_ = json.NewEncoder(w).Encode(h)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", headers)
	_ = http.ListenAndServe(":8000", r)
}
