package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

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
	srv := http.HandlerFunc(responseFunction)
	srv.ServeHTTP(rr, req)
	return rr
}

func checkResponse(t *testing.T, resp *httptest.ResponseRecorder, expected int) {
	if expected != resp.Code {
		t.Errorf("Expected response code %d and got %d.\n. Response body: %s\n", expected, resp.Code, resp.Body.String())
	}
}
