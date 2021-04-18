package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

func TestCleanDir(t *testing.T) {
	hashanddelete()
}

func TestHash(t *testing.T) {
	hashfile("main.go")
	hashfile("nothere")
}

func TestHello(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, Hello)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestHeaders(t *testing.T) {
	r, _ := http.NewRequest("GET", "/headers", nil)
	r.Header.Add("hello", "world")
	rr := executeRequest(r, EchoHeaders)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestAllowedMethod(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, AllowedMethod(Hello, "POST"))
	checkResponseCode(t, http.StatusMethodNotAllowed, rr.Code)
	r, _ = http.NewRequest("GET", "/", nil)
	rr = executeRequest(r, AllowedMethod(Hello, "GET,POST"))
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestRedirect(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/redirect/302", nil)
	rr := executeVarsRequest("/redirect/{code}", r, Redirect)
	checkResponseCode(t, http.StatusFound, rr.Code)
	r, _ = http.NewRequest("GET", "/redirect/500", nil)
	rr = executeVarsRequest("/redirect/{code}", r, Redirect)
	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	r, _ = http.NewRequest("GET", "/redirect/hello", nil)
	rr = executeVarsRequest("/redirect/{code}", r, Redirect)
	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	r, _ = http.NewRequest("GET", "/redirect/", nil)
	rr = executeRequest(r, Redirect)
	if !(rr.Code <= 307 && rr.Code >= 300) {
		t.Errorf("Expected redirect code between 300 and 307 got %d\n", rr.Code)
	}
}

func TestStatusCodeGood(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status", nil)
	rr := executeRequest(r, StatusCode)
	checkResponseCode(t, http.StatusOK, rr.Code)

}

func TestStatusCode404(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/404", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponseCode(t, http.StatusNotFound, rr.Code)
}

func TestStatusCode500(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/500", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
}

func TestStatusCodeFloat(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/5.1", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestStatusCodeString(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/fail", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestStatusCodeExtra(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/700", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponseCode(t, 700, rr.Code)
}

func executeVarsRequest(path string, req *http.Request, responseFunction func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, responseFunction)
	router.ServeHTTP(rr, req)
	return rr
}

func executeRequest(req *http.Request, responseFunction func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	http := http.HandlerFunc(responseFunction)
	http.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
