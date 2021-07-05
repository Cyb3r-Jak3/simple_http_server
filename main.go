package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

var host string
var port string

func hashfile(filename string) {
	f, err := os.Open(filename) // #nosec
	if err != nil {
		log.Printf("Couldn't open %s. Error reason %s\n", filename, err.Error())
		return
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Printf("Couldn't hash %s. Error reason %s\n", filename, err.Error())
	}
	fmt.Printf("Hash of %s is %x\n", filename, h.Sum(nil))
	if err := f.Close(); err != nil {
		log.Printf("Error closing file: %s\n", err)
	}
}

func hashanddelete() {
	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dir {
		hashfile(path.Join([]string{dirName, d.Name()}...))
		log.Printf("Removing %s\n", d.Name())
		if err := os.RemoveAll(path.Join([]string{dirName, d.Name()}...)); err != nil {
			log.Printf("Error deleting file %s\n", err)
		}
	}
}

func cleardir() {
	if err := os.Mkdir(dirName, 0200); err != nil {
		log.Printf("Error creating directory %s\n", err)
	}

	for {
		hashanddelete()
		time.Sleep(cleanSeconds * time.Second)
	}
}

func init() {
	host = common.GetEnv("HOST", "")
	port = common.GetEnv("PORT", "8090")
}

func main() {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	go cleardir()
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/headers", EchoHeaders)
	r.HandleFunc("/post/json", common.AllowedMethod(PostJSON, "POST"))
	r.HandleFunc("/post/file/form", common.AllowedMethod(PostFormFile, "POST"))
	r.HandleFunc("/post/file/{name}", common.AllowedMethod(PostFile, "POST"))
	r.HandleFunc("/get/json", common.AllowedMethod(GetJSON, "GET"))
	r.HandleFunc("/get/json/{rows}", common.AllowedMethod(GetJSON, "GET"))
	r.HandleFunc("/get/image", common.AllowedMethod(GetImage, "GET"))
	r.HandleFunc("/get/image/{type}", common.AllowedMethod(GetImage, "GET"))
	r.HandleFunc("/get/image/{type}/{height}", common.AllowedMethod(GetImage, "GET"))
	r.HandleFunc("/get/image/{type}/{height}/{width}", common.AllowedMethod(GetImage, "GET"))
	r.HandleFunc("/get/uuid", common.AllowedMethod(GetUUID, "GET"))
	r.HandleFunc("/get/ipv4", common.AllowedMethod(GetIPv4, "GET"))
	r.HandleFunc("/get/ipv6", common.AllowedMethod(GetIPv6, "GET"))
	r.HandleFunc("/get/base64", common.AllowedMethod(GetBase64, "GET"))
	r.HandleFunc("/get/xml", common.AllowedMethod(GetXML, "GET"))
	r.HandleFunc("/get/xml/{rows}", common.AllowedMethod(GetXML, "GET"))
	r.HandleFunc("/get/csv", common.AllowedMethod(GetCSV, "GET"))
	r.HandleFunc("/get/csv/{rows}", common.AllowedMethod(GetCSV, "GET"))
	r.HandleFunc("/cookies/get", common.AllowedMethod(GetCookies, "GET,POST"))
	r.HandleFunc("/cookies/set/{name}/{value}", common.AllowedMethod(SetCookie, "GET,POST"))
	r.HandleFunc("/cookies/clear", common.AllowedMethod(ClearCookies, "GET,POST"))
	r.HandleFunc("/status", StatusCode)
	r.HandleFunc("/status/{code}", StatusCode)
	r.HandleFunc("/redirect", Redirect)
	r.HandleFunc("/redirect/{code}", Redirect)
	r.HandleFunc("/auth/basic/{username}/{password}", DynamicAuth)
	r.HandleFunc("/auth/basic/bad", BasicAuth(Hello, "admin", "admin"))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r)
	if err != nil {
		log.Fatal(err)
	}
}
