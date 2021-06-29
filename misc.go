package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/gorilla/mux"
)

// Hello is a simple hello function
func Hello(w http.ResponseWriter, _ *http.Request) { common.StringResponse(w, "Hello") }

// EchoHeaders returns the headers of the request
func EchoHeaders(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// StatusCode returns either the given code or a random one
func StatusCode(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if vars["code"] == "" {
		common.StringResponse(w, "Ok")
	} else {
		i, err := strconv.Atoi(vars["code"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(i)
		if _, err := w.Write([]byte(fmt.Sprintf("Status Code: %d", i))); err != nil {
			log.Printf("Error writing status code response: %s\n", err)
		}
	}
}
