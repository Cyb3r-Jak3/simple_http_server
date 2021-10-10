package main

import (
	"crypto/subtle"
	"log"
	"net/http"

	"github.com/Cyb3r-Jak3/common/v2"
	"github.com/gorilla/mux"
)

const defaultAuth = "admin"

// DynamicAuth does basic authorization for given username and password
func DynamicAuth(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var username, password string
	if vars["username"] == "" {
		username = defaultAuth
		password = defaultAuth
	} else {
		username = vars["username"]
		password = vars["password"]
	}
	user, pass, ok := req.BasicAuth()

	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
		w.Header().Set("WWW-Authenticate", `Basic realm="Simple HTTP Server"`)
		w.WriteHeader(401)
		if _, err := w.Write([]byte("Unauthorized.\n")); err != nil {
			log.Printf("Error writing authorized response: %s\n", err)
		}
		return
	}
	common.StringResponse(w, "Authenticated")
}
