package main

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestGETJSON(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("POST", "/get/json", nil)
	rr := executeRequest(r, GetJSON)
	checkResponseCode(t, http.StatusMethodNotAllowed, rr.Code)
	r, _ = http.NewRequest("GET", "/get/json", nil)
	rr = executeRequest(r, GetJSON)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("GET", "/get/json/5", nil)
	rr = executeVarsRequest("/get/json/{rows}", r, GetJSON)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("GET", "/get/json/hello", nil)
	rr = executeVarsRequest("/get/json/{rows}", r, GetJSON)
	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}

func TestGetImagePNG(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/image/png", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("GET", "/get/image", nil)
	rr = executeRequest(r, GetImage)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("GET", "/get/image/png/200", nil)
	rr = executeVarsRequest("/get/image/{type}/{height}", r, GetImage)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("GET", "/get/image/png/200/200", nil)
	rr = executeVarsRequest("/get/image/{type}/{height}/{width}", r, GetImage)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("GET", "/get/image/200/200", nil)
	rr = executeVarsRequest("/get/image/{type}/{height}", r, GetImage)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestBadImage(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("POST", "/get/image/png", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponseCode(t, http.StatusMethodNotAllowed, rr.Code)

}

func TestGetImageJPG(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/image/jpg", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponseCode(t, http.StatusOK, rr.Code)

}
func TestGetImageURL(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/image/url", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestGetImageFake(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/image/fake", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponseCode(t, http.StatusOK, rr.Code)
}

func TestGetUUID(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/uuid", nil)
	rr := executeRequest(r, GetUUID)
	// Using regex because I don't want to import a uuid package for testing
	reg := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	if !reg.MatchString(rr.Body.String()) {
		t.Error("UUID was not returned")
	}
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("POST", "/get/uuid", nil)
	rr = executeRequest(r, GetUUID)
	checkResponseCode(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestGetIPv4(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/ipv4", nil)
	rr := executeRequest(r, GetIPv4)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("POST", "/get/ipv4", nil)
	rr = executeRequest(r, GetIPv4)
	checkResponseCode(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestGetIPv6(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/ipv6", nil)
	rr := executeRequest(r, GetIPv6)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("POST", "/get/ipv6", nil)
	rr = executeRequest(r, GetIPv6)
	checkResponseCode(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestGetBase64(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/base64", nil)
	rr := executeRequest(r, GetBase64)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("POST", "/get/base64", nil)
	rr = executeRequest(r, GetBase64)
	checkResponseCode(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestGETXML(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("POST", "/get/xml", nil)
	rr := executeRequest(r, GetXML)
	checkResponseCode(t, http.StatusMethodNotAllowed, rr.Code)
	r, _ = http.NewRequest("GET", "/get/xml", nil)
	rr = executeRequest(r, GetXML)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("GET", "/get/xml/5", nil)
	rr = executeVarsRequest("/get/xml/{rows}", r, GetXML)
	checkResponseCode(t, http.StatusOK, rr.Code)
	r, _ = http.NewRequest("GET", "/get/xml/hello", nil)
	rr = executeVarsRequest("/get/xml/{rows}", r, GetXML)
	checkResponseCode(t, http.StatusBadRequest, rr.Code)
}
