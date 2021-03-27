package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

func CheckMethod(method string, req *http.Request) bool {
	return method == req.Method
}

func clear_dir() {
	os.Mkdir(dir_name, 0200)

	for {
		dir, err := ioutil.ReadDir(dir_name)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range dir {
			log.Printf("Removing %s", d.Name())
			os.RemoveAll(path.Join([]string{"tmp", d.Name()}...))
		}
		time.Sleep(clean_seconds * time.Second)
	}
}

func main() {
	Faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(Faker)
	go clear_dir()
	r := mux.NewRouter()
	r.HandleFunc("/", Hello).Methods("GET")
	r.HandleFunc("/headers", EchoHeaders)
	r.HandleFunc("/post/json", PostJSON)
	r.HandleFunc("/post/form/file", PostFormFile)
	r.HandleFunc("/post/file/{name}", PostFile)
	r.HandleFunc("/get/json", GetJson)
	r.HandleFunc("/get/json/{row}", GetJson)
	r.HandleFunc("/get/image", GetImage)
	r.HandleFunc("/get/image/{type}", GetImage)
	r.HandleFunc("/get/image/{type}/{height}", GetImage)
	r.HandleFunc("/get/image/{type}/{height}/{width}", GetImage)
	r.HandleFunc("/cookies/set/{name}/{value}", SetCookie)
	r.HandleFunc("/cookies/get", GetCookies)
	r.HandleFunc("/cookies/clear", ClearCookies)

	err := http.ListenAndServe(":8090", r)
	if err != nil {
		log.Fatal(err)
	}
}
