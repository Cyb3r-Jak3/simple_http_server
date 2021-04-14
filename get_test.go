package main

import (
	"net/http"
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
