package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"net/http"
	"testing"
)

func TestRedirect(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/redirect/302", nil)
	rr := executeVarsRequest("/redirect/{code}", r, Redirect)
	checkResponse(t, rr, http.StatusFound)
	r, _ = http.NewRequest("GET", "/redirect/500", nil)
	rr = executeVarsRequest("/redirect/{code}", r, Redirect)
	checkResponse(t, rr, http.StatusBadRequest)
	r, _ = http.NewRequest("GET", "/redirect/hello", nil)
	rr = executeVarsRequest("/redirect/{code}", r, Redirect)
	checkResponse(t, rr, http.StatusBadRequest)
	r, _ = http.NewRequest("GET", "/redirect/", nil)
	rr = executeRequest(r, Redirect)
	if !(rr.Code <= 307 && rr.Code >= 300) {
		t.Errorf("Expected redirect code between 300 and 307 got %d\n", rr.Code)
	}
}