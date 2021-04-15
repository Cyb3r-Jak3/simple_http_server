package main

import (
	"crypto/subtle"
	"net/http"

	"github.com/gorilla/mux"
)

// DynamicAuth does basic authorization for given username and password
func DynamicAuth(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if vars["username"] == "" || vars["password"] == "" {
		http.Error(w, "User and password is required.", http.StatusBadRequest)
	}
	user, pass, ok := req.BasicAuth()

	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(vars["username"])) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(vars["password"])) != 1 {
		w.Header().Set("WWW-Authenticate", `Basic realm="Simple HTTP Server"`)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized.\n"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Authenticated"))
}

// BasicAuth handles HTTP basic auth
func BasicAuth(handler http.HandlerFunc, username, password string) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		user, pass, ok := req.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="Simple HTTP Server"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}

		handler(w, req)
	}
}
