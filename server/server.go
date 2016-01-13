package server

import (
	"net/http"
	"fmt"
	"github.com/doojin/doc-to-html-service/service/dirscanner"
	"github.com/gorilla/mux"
)

var Scanner dirscanner.DirScannerInterface = new(dirscanner.DirScanner)

// Start starts the web server which is responsible for receiving
// requests for converting
func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", process).Methods("POST")

	fmt.Println("Server was started")

	http.Handle("/", r)
	http.ListenAndServe(":8030", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	dir := r.FormValue("dir")
	files := Scanner.ScanDir(dir)
	fmt.Println(files)
}