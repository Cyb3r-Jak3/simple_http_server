package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/gorilla/mux"
)

// PostJSON echos JSON back that it was sent
func PostJSON(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, maxUploadSize*1024*1024)
	if req.Body == http.NoBody || req.ContentLength == 0 {
		http.Error(w, "JSON body required", http.StatusBadRequest)
		return
	}
	out, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	common.JSONResponse(w, out)
}

// PostFormFile saves a file that upload as a form-data/multipart request
func PostFormFile(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, maxUploadSize*1024*1024)
	file, handler, err := req.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	f, err := os.OpenFile(filepath.Join(dirName, handler.Filename), os.O_WRONLY|os.O_CREATE, 0200)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := io.Copy(f, file); err != nil {
		log.Printf("Error with io.Copy: %s\n", err)
		http.Error(w, "Error with copying file", http.StatusInternalServerError)
		return
	}
	common.StringResponse(w, "File uploaded")
	if err := file.Close(); err != nil {
		log.Printf("Error closing file: %s\n", err)
	}
	if err := f.Close(); err != nil {
		log.Printf("Error closing f: %s\n", err)
	}
}

//PostFile saves a file that is posted in a request body
func PostFile(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, maxUploadSize*1024*1024)
	vars := mux.Vars(req)
	f, err := os.OpenFile(filepath.Join(dirName, vars["name"]), os.O_WRONLY|os.O_CREATE, 0200)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(f, req.Body); err != nil {
		log.Printf("Error with io.Copy: %s\n", err)
		http.Error(w, "Error with copying file", http.StatusInternalServerError)
		return
	}
	common.StringResponse(w, "File uploaded")
	if err := f.Close(); err != nil {
		log.Printf("Error closing f: %s\n", err)
	}
}
