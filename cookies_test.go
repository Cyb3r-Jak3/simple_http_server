package main

import (
	"net/http"
	"testing"
	"time"
)

func TestGetCookies(t *testing.T) {
	r, _ := http.NewRequest("GET", "/cookies/get", nil)
	r.AddCookie(
		&http.Cookie{
			Name:     "testCookie",
			Value:    "Some Data",
			HttpOnly: true,
			Domain:   cookieDomain,
			Expires:  time.Now(),
			Secure:   true,
		},
	)
	rr := executeRequest(r, GetCookies)
	checkResponse(t, rr, http.StatusOK)
}

func TestSetCookies(t *testing.T) {
	r, _ := http.NewRequest("GET", "/cookies/set/hello/world", nil)
	rr := executeVarsRequest("/cookies/set/{name}/{value}", r, SetCookie)
	if len(rr.Result().Cookies()) == 0 {
		t.Errorf("expected at least one cookie got %d cookies", len(rr.Result().Cookies()))
	}
	checkResponse(t, rr, http.StatusOK)
}

func TestClearCookies(t *testing.T) {
	r, _ := http.NewRequest("GET", "/cookies/clear", nil)
	r.AddCookie(
		&http.Cookie{
			Name:     "testCookie",
			Value:    "Some Data",
			HttpOnly: true,
			// Domain:   cookie_domain,
			Expires: time.Now(),
			Secure:  true,
		},
	)
	rr := executeRequest(r, ClearCookies)
	checkResponse(t, rr, http.StatusOK)
}
