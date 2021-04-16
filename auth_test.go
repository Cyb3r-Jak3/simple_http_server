package main

import (
	"net/http"
	"testing"
)

func TestBadAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/auth/basic/bad", nil)
	rr := executeRequest(r, BasicAuth(Hello, "user", "user"))
	checkResponseCode(t, http.StatusUnauthorized, rr.Code)
}

func TestGoodAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/auth/basic/bad", nil)
	r.SetBasicAuth("admin", "admin")
	rr := executeRequest(r, BasicAuth(Hello, "admin", "admin"))
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestBadDynamicAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/auth/basic/user/password", nil)
	r.SetBasicAuth("user", "failed")
	rr := executeVarsRequest("/auth/basic/{user}/{password}", r, DynamicAuth)
	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	r, _ = http.NewRequest("GET", "/auth/basic/user/bad", nil)
	r.SetBasicAuth("user", "failed")
	rr = executeVarsRequest("/auth/basic/{username}/{password}", r, DynamicAuth)
	checkResponseCode(t, http.StatusUnauthorized, rr.Code)
}

func TestGoodDynamicAuth(t *testing.T) {
	r, _ := http.NewRequest("GET", "/auth/basic/user/allowed", nil)
	r.SetBasicAuth("user", "allowed")
	rr := executeVarsRequest("/auth/basic/{username}/{password}", r, DynamicAuth)
	checkResponseCode(t, http.StatusOK, rr.Code)
}
