package server

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", process).Methods("POST")

	fmt.Println("Server was started")

	http.Handle("/", r)
	http.ListenAndServe(":8030", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request was received")
}