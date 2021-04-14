package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

// PostJSON echos JSON back that it was sent
func PostJSON(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("POST", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	req.Body = http.MaxBytesReader(w, req.Body, maxUploadSize*1024*1024)
	if req.Body == http.NoBody || req.ContentLength == 0 {
		http.Error(w, "JSON body required", http.StatusBadRequest)
		return
	}
	out, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// PostFormFile saves a file that upload as a form-data/multipart request
func PostFormFile(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("POST", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	req.Body = http.MaxBytesReader(w, req.Body, maxUploadSize*1024*1024)
	file, handler, err := req.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()
	f, err := os.OpenFile(filepath.Join(dirName, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintln(w, "Done")
}

//PostFile saves a file that is posted in a request body
func PostFile(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("POST", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	req.Body = http.MaxBytesReader(w, req.Body, maxUploadSize*1024*1024)
	vars := mux.Vars(req)
	f, err := os.OpenFile(filepath.Join(dirName, vars["name"]), os.O_WRONLY|os.O_CREATE, 0200)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(f, req.Body)
	fmt.Fprintln(w, "Done")
}
