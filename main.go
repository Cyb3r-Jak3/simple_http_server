package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

// AllowedMethod is a decorator to get methods
func AllowedMethod(handler http.HandlerFunc, methods string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		for _, b := range strings.Split(methods, ",") {
			if b == req.Method {
				handler(w, req)
			}
		}
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

// StringResponse writes a http response as a string
func StringResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("Error with writing string response: %s\n", err)
	}
}

// JSONResponse writes a http response as JSON
func JSONResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error with writing JSON response: %s\n", err)
	}
}

// ContentResponse writes a http response with a given content type
func ContentResponse(w http.ResponseWriter, contentType string, response []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error with writing content type '%s' response: %s\n", contentType, err)
	}
}

// GenerateRandInt securely generate a random int64
func GenerateRandInt(x int) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(x)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()), nil
}

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
	dir, err := os.ReadDir(dirName)
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

func main() {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	go cleardir()
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/headers", EchoHeaders)
	r.HandleFunc("/post/json", AllowedMethod(PostJSON, "POST"))
	r.HandleFunc("/post/file/form", AllowedMethod(PostFormFile, "POST"))
	r.HandleFunc("/post/file/{name}", AllowedMethod(PostFile, "POST"))
	r.HandleFunc("/get/json", AllowedMethod(GetJSON, "GET"))
	r.HandleFunc("/get/json/{rows}", AllowedMethod(GetJSON, "GET"))
	r.HandleFunc("/get/image", AllowedMethod(GetImage, "GET"))
	r.HandleFunc("/get/image/{type}", AllowedMethod(GetImage, "GET"))
	r.HandleFunc("/get/image/{type}/{height}", AllowedMethod(GetImage, "GET"))
	r.HandleFunc("/get/image/{type}/{height}/{width}", AllowedMethod(GetImage, "GET"))
	r.HandleFunc("/get/uuid", AllowedMethod(GetUUID, "GET"))
	r.HandleFunc("/get/ipv4", AllowedMethod(GetIPv4, "GET"))
	r.HandleFunc("/get/ipv6", AllowedMethod(GetIPv6, "GET"))
	r.HandleFunc("/get/base64", AllowedMethod(GetBase64, "GET"))
	r.HandleFunc("/get/xml", AllowedMethod(GetXML, "GET"))
	r.HandleFunc("/get/xml/{rows}", AllowedMethod(GetXML, "GET"))
	r.HandleFunc("/cookies/get", AllowedMethod(GetCookies, "GET,POST"))
	r.HandleFunc("/cookies/set/{name}/{value}", AllowedMethod(SetCookie, "GET,POST"))
	r.HandleFunc("/cookies/clear", AllowedMethod(ClearCookies, "GET,POST"))
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
