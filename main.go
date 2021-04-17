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

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

// CheckMethod checks to make sure a request has an allowed method
func CheckMethod(method string, req *http.Request) bool {
	return method == req.Method
}

func hashfile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Couldn't open %s. Error reason %s", filename, err.Error())
		return
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Printf("Couldn't hash %s. Error reason %s", filename, err.Error())
	}
	fmt.Printf("Hash of %s is %x\n", filename, h.Sum(nil))
}

func hashanddelete() {
	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dir {
		hashfile(path.Join([]string{dirName, d.Name()}...))
		log.Printf("Removing %s", d.Name())
		os.RemoveAll(path.Join([]string{dirName, d.Name()}...))
	}
}

func cleardir() {
	os.Mkdir(dirName, 0200)

	for {
		hashanddelete()
		time.Sleep(cleanSeconds * time.Second)
	}
}

func main() {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	go cleardir()
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/headers", EchoHeaders)
	r.HandleFunc("/post/json", PostJSON)
	r.HandleFunc("/post/file/form", PostFormFile)
	r.HandleFunc("/post/file/{name}", PostFile)
	r.HandleFunc("/get/json", GetJSON)
	r.HandleFunc("/get/json/{rows}", GetJSON)
	r.HandleFunc("/get/image", GetImage)
	r.HandleFunc("/get/image/{type}", GetImage)
	r.HandleFunc("/get/image/{type}/{height}", GetImage)
	r.HandleFunc("/get/image/{type}/{height}/{width}", GetImage)
	r.HandleFunc("/get/uuid", GetUUID)
	r.HandleFunc("/get/ipv4", GetIPv4)
	r.HandleFunc("/get/ipv6", GetIPv6)
	r.HandleFunc("/get/base64", GetBase64)
	r.HandleFunc("/get/xml", GetXML)
	r.HandleFunc("/get/xml/{rows}", GetXML)
	r.HandleFunc("/cookies/get", GetCookies)
	r.HandleFunc("/cookies/set/{name}/{value}", SetCookie)
	r.HandleFunc("/cookies/clear", ClearCookies)
	r.HandleFunc("/status", StatusCode)
	r.HandleFunc("/status/{code}", StatusCode)
	r.HandleFunc("/redirect", Redirect)
	r.HandleFunc("/redirect/{code}", Redirect)
	r.HandleFunc("/auth/basic/{username}/{password}", DynamicAuth)
	r.HandleFunc("/auth/basic/bad", BasicAuth(Hello, "admin", "admin"))

	err := http.ListenAndServe(":8090", r)
	if err != nil {
		log.Fatal(err)
	}
}
