package main

import (
	"net/http"
	"testing"
)

func TestBadAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/auth/basic/bad", nil)
	rr := executeRequest(r, BasicAuth(Hello, "user", "user"))
	checkResponse(t, rr, http.StatusUnauthorized)
}

func TestGoodAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/auth/basic/bad", nil)
	r.SetBasicAuth("admin", "admin")
	rr := executeRequest(r, BasicAuth(Hello, "admin", "admin"))
	checkResponse(t, rr, http.StatusOK)
}

func TestBadDynamicAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/auth/basic/user/password", nil)
	r.SetBasicAuth("user", "failed")
	rr := executeVarsRequest("/auth/basic/{user}/{password}", r, DynamicAuth)
	checkResponse(t, rr, http.StatusBadRequest)
	r, _ = http.NewRequest("GET", "/auth/basic/user/bad", nil)
	r.SetBasicAuth("user", "failed")
	rr = executeVarsRequest("/auth/basic/{username}/{password}", r, DynamicAuth)
	checkResponse(t, rr, http.StatusUnauthorized)
}

func TestGoodDynamicAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/auth/basic/user/allowed", nil)
	r.SetBasicAuth("user", "allowed")
	rr := executeVarsRequest("/auth/basic/{username}/{password}", r, DynamicAuth)
	checkResponse(t, rr, http.StatusOK)
}
