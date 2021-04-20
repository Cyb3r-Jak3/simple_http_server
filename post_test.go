package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestPostJSON(t *testing.T) {
	r, _ := http.NewRequest("POST", "/post/json", bytes.NewBuffer([]byte(`{"hello":"world"}`)))
	r.Header.Set("Content-Type", "application/json")
	rr := executeRequest(r, PostJSON)
	checkResponse(t, rr, http.StatusOK)
}

func TestPostBadJSON(t *testing.T) {
	r, _ := http.NewRequest("POST", "/post/json", nil)
	r.Header.Set("Content-Type", "application/json")
	rr := executeRequest(r, PostJSON)
	checkResponse(t, rr, http.StatusBadRequest)
}

func TestPostFormFile(t *testing.T) {
	file, _ := os.Open("main.go")
	fileContents, _ := io.ReadAll(file)
	file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "main")
	part.Write(fileContents)
	r, _ := http.NewRequest("POST", "/post/form/file", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	writer.Close()
	rr := executeRequest(r, PostFormFile)
	checkResponse(t, rr, http.StatusOK)
	hashanddelete()
}

func TestPostFile(t *testing.T) {
	file, _ := os.Open("main.go")
	defer file.Close()
	r, _ := http.NewRequest("POST", "/post/file/main", file)
	r.Header.Add("Content-Type", "binary/octet-stream")
	rr := executeVarsRequest("/post/file/{name}", r, PostFile)
	checkResponse(t, rr, http.StatusOK)
}
