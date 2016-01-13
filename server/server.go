package server

import (
	"net/http"
	"fmt"
)

func Start() {
	http.HandleFunc("/", process)
	fmt.Println("Server was started")
	http.ListenAndServe(":8030", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request was received")
}