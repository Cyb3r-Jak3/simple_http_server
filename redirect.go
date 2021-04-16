package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Redirect will redirect users to the given domain with either a random code or a supplied one
func Redirect(w http.ResponseWriter, req *http.Request) {
	var redirectCode int
	var err error
	vars := mux.Vars(req)
	if vars["code"] == "" {
		redirectCode = Faker.Number(300, 307)
	} else {
		redirectCode, err = strconv.Atoi(vars["code"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	if redirectCode <= 307 && redirectCode >= 300 {
		http.Redirect(w, req, redirectURL, redirectCode)
	} else {
		http.Error(w, "Not a valid redirect HTTP code", http.StatusBadRequest)
	}

}
