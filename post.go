package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func PostJSON(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("POST", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	buf := new(strings.Builder)
	_, err := io.Copy(buf, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())
}

func PostFormFile(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("POST", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	req.Body = http.MaxBytesReader(w, req.Body, max_upload_size*1024*1024)
	file, handler, err := req.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()
	f, err := os.OpenFile(filepath.Join(dir_name, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprint(w, "Done")
}

func PostFile(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("POST", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	vars := mux.Vars(req)
	f, err := os.OpenFile(filepath.Join(dir_name, vars["name"]), os.O_WRONLY|os.O_CREATE, 0200)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(f, req.Body)
	fmt.Fprint(w, "Done")

}
