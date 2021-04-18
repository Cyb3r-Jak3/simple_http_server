package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Hello is a simple hello function
func Hello(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello"))
}

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
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	} else {
		i, err := strconv.Atoi(vars["code"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(i)
		w.Write([]byte(fmt.Sprintf("Status Code: %d", i)))
	}
}
