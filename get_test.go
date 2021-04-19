package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestGETJSON(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/json", nil)
	rr := executeRequest(r, GetJSON)
	checkResponse(t, rr, http.StatusOK)
	r, _ = http.NewRequest("GET", "/get/json/5", nil)
	rr = executeVarsRequest("/get/json/{rows}", r, GetJSON)
	checkResponse(t, rr, http.StatusOK)
	r, _ = http.NewRequest("GET", "/get/json/hello", nil)
	rr = executeVarsRequest("/get/json/{rows}", r, GetJSON)
	checkResponse(t, rr, http.StatusBadRequest)
}

func TestGetImagePNG(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/image/png", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponse(t, rr, http.StatusOK)
	r, _ = http.NewRequest("GET", "/get/image", nil)
	rr = executeRequest(r, GetImage)
	checkResponse(t, rr, http.StatusOK)
	r, _ = http.NewRequest("GET", "/get/image/png/200", nil)
	rr = executeVarsRequest("/get/image/{type}/{height}", r, GetImage)
	checkResponse(t, rr, http.StatusOK)
	r, _ = http.NewRequest("GET", "/get/image/png/200/200", nil)
	rr = executeVarsRequest("/get/image/{type}/{height}/{width}", r, GetImage)
	checkResponse(t, rr, http.StatusOK)
	r, _ = http.NewRequest("GET", "/get/image/200/200", nil)
	rr = executeVarsRequest("/get/image/{type}/{height}", r, GetImage)
	checkResponse(t, rr, http.StatusOK)
}

func TestGetImageJPG(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/image/jpg", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponse(t, rr, http.StatusOK)

}
func TestGetImageURL(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/image/url", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponse(t, rr, http.StatusOK)
}

func TestGetImageFake(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/image/fake", nil)
	rr := executeVarsRequest("/get/image/{type}", r, GetImage)
	checkResponse(t, rr, http.StatusOK)
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
	checkResponse(t, rr, http.StatusOK)
}

func TestGetIPv4(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/ipv4", nil)
	rr := executeRequest(r, GetIPv4)
	checkResponse(t, rr, http.StatusOK)
}

func TestGetIPv6(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/ipv6", nil)
	rr := executeRequest(r, GetIPv6)
	checkResponse(t, rr, http.StatusOK)
}

func TestGetBase64(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/base64", nil)
	rr := executeRequest(r, GetBase64)
	checkResponse(t, rr, http.StatusOK)
	_, err := base64.URLEncoding.DecodeString(rr.Body.String())
	if err != nil {
		t.Errorf("Expected to decode base64 got %s", err.Error())
	}
}

func TestGETXML(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	fmt.Println("Use Server rows")
	r, _ := http.NewRequest("GET", "/get/xml", nil)
	rr := executeRequest(r, GetXML)
	checkResponse(t, rr, http.StatusOK)
	fmt.Println("Sending rows")
	r, _ = http.NewRequest("GET", "/get/xml/5", nil)
	rr = executeVarsRequest("/get/xml/{rows}", r, GetXML)
	checkResponse(t, rr, http.StatusOK)
	fmt.Printf("String row")
	r, _ = http.NewRequest("GET", "/get/xml/hello", nil)
	rr = executeVarsRequest("/get/xml/{rows}", r, GetXML)
	checkResponse(t, rr, http.StatusBadRequest)
}

func TestGetCSV(t *testing.T) {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	r, _ := http.NewRequest("GET", "/get/csv", nil)
	rr := executeRequest(r, GetCSV)
	checkResponse(t, rr, http.StatusOK)
	r, _ = http.NewRequest("GET", "/get/csv/5", nil)
	rr = executeVarsRequest("/get/csv/{rows}", r, GetCSV)
	checkResponse(t, rr, http.StatusOK)
	r, _ = http.NewRequest("GET", "/get/csv/hello", nil)
	rr = executeVarsRequest("/get/csv/{rows}", r, GetCSV)
	checkResponse(t, rr, http.StatusBadRequest)
}
