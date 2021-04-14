package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
		fmt.Fprintf(w, "OK")
	} else {
		i, err := strconv.Atoi(vars["code"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		http.Error(w, "", i)
	}
}
