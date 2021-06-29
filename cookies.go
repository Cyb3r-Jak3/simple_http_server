package main

import (
	"fmt"
	"net/http"
	"time"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/gorilla/mux"
)

// GetCookies gets and prints all cookies from a request
func GetCookies(w http.ResponseWriter, req *http.Request) {
	for _, cookie := range req.Cookies() {
		fmt.Fprintf(w, "%v\n", cookie)
	}
}

//SetCookie sets a new cookie
func SetCookie(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	experation := time.Now().Add(24 * time.Hour)
	cookie := &http.Cookie{
		Name:     vars["name"],
		Value:    vars["value"],
		HttpOnly: true,
		Domain:   cookieDomain,
		Expires:  experation,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
}

//ClearCookies clears all the cookies from a request
func ClearCookies(w http.ResponseWriter, req *http.Request) {
	for _, cookie := range req.Cookies() {
		clearedCookie, _ := req.Cookie(cookie.Name)
		clearedCookie.MaxAge = -1
		http.SetCookie(w, clearedCookie)
	}
	common.StringResponse(w, "All Cookies Should Be Cleared")
}
