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
	checkResponse(t, rr, http.StatusOK)
}

func TestHeaders(t *testing.T) {
	r, _ := http.NewRequest("GET", "/headers", nil)
	r.Header.Add("hello", "world")
	rr := executeRequest(r, EchoHeaders)
	checkResponse(t, rr, http.StatusOK)
}

func TestAllowedMethod(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, AllowedMethod(Hello, "POST"))
	checkResponse(t, rr, http.StatusMethodNotAllowed)
	r, _ = http.NewRequest("GET", "/", nil)
	rr = executeRequest(r, AllowedMethod(Hello, "GET,POST"))
	checkResponse(t, rr, http.StatusOK)
}

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

func TestStatusCodeGood(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status", nil)
	rr := executeRequest(r, StatusCode)
	checkResponse(t, rr, http.StatusOK)

}

func TestStatusCode404(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/404", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponse(t, rr, http.StatusNotFound)
}

func TestStatusCode500(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/500", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponse(t, rr, http.StatusInternalServerError)
}

func TestStatusCodeFloat(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/5.1", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponse(t, rr, http.StatusBadRequest)
}

func TestStatusCodeString(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/fail", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponse(t, rr, http.StatusBadRequest)
}

func TestStatusCodeExtra(t *testing.T) {
	r, _ := http.NewRequest("GET", "/status/700", nil)
	rr := executeVarsRequest("/status/{code}", r, StatusCode)
	checkResponse(t, rr, 700)
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

func checkResponse(t *testing.T, resp *httptest.ResponseRecorder, expected int) {
	if expected != resp.Code {
		t.Errorf("Expected response code %d. Got %d\n. Response body %s\n", expected, resp.Code, resp.Body.String())
	}
}
