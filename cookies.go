package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetCookies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for _, cookie := range req.Cookies() {
		fmt.Fprintf(w, "%v\n", cookie)
	}
}

func SetCookie(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	experation := time.Now().Add(24 * time.Hour)
	cookie := &http.Cookie{
		Name:     vars["name"],
		Value:    vars["value"],
		HttpOnly: true,
		// Domain:   cookie_domain,
		Expires: experation,
		Secure:  true,
	}
	http.SetCookie(w, cookie)
}

func ClearCookies(w http.ResponseWriter, req *http.Request) {
	for _, cookie := range req.Cookies() {
		clear_cookie, _ := req.Cookie(cookie.Name)
		clear_cookie.MaxAge = -1
		fmt.Println(clear_cookie.MaxAge)
		http.SetCookie(w, clear_cookie)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "All Cookies Should Be Cleared")
}
