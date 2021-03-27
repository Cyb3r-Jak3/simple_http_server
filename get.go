package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func GetJson(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("GET", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	vars := mux.Vars(req)
	var rowCount int
	if vars["row"] == "" {
		rowCount = rand.Intn(row_count)
	} else {
		rowCount, _ = strconv.Atoi(vars["row"])
	}
	garbage, err := Faker.JSON(&gofakeit.JSONOptions{
		Type:     "array",
		RowCount: rowCount,
		Fields: []gofakeit.Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "email", Function: "email"},
			{Name: "tag", Function: "gamertag"},
		},
		Indent: true,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(garbage))
}

func downloadImage(url string, format string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	respImage, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if format == "png" {
		//Copied from https://github.com/tizz98/comix/blob/master/app/img.go
		respImage, err := jpeg.Decode(bytes.NewReader(respImage))
		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		if err := png.Encode(buf, respImage); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
	return respImage, nil

}

func isvalidformat(format string) bool {
	switch format {
	case
		"jpg",
		"png",
		"url":
		return true
	}
	return false
}

func GeneratorImageURL(vars map[string]string) string {
	size_options := [7][2]string{
		{"16", "16"},
		{"32", "32"},
		{"64", "64"},
		{"480", "360"},
		{"640", "480"},
		{"1080", "720"},
		{"1920", "1080"},
	}
	if vars["width"] == "" {
		if vars["height"] == "" {
			s := rand.NewSource(time.Now().Unix())
			r := rand.New(s)
			randomIndex := r.Intn(len(size_options))
			pick := size_options[randomIndex]
			return fmt.Sprint(base_image_url, fmt.Sprintf("%s/%s.jpg", pick[0], pick[1]))
		} else {
			if isvalidformat(vars["type"]) {
				return fmt.Sprint(base_image_url, fmt.Sprintf("%s.jpg", vars["height"]))
			} else {
				return fmt.Sprint(base_image_url, fmt.Sprintf("%s/%s.jpg", vars["type"], vars["height"]))
			}
		}

	}
	return fmt.Sprint(base_image_url, fmt.Sprintf("%s/%s.jpg", vars["height"], vars["width"]))

}

func GetImage(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("GET", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	var Image []byte
	var ImageErr error
	var imageType, imageURL string
	vars := mux.Vars(req)
	if vars["type"] == "" {
		image_types := make([]string, 0)
		image_types = append(image_types,
			"png",
			"jpg",
			"url",
		)
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		imageType = image_types[r.Intn(len(image_types))]
	} else {
		imageType = vars["type"]
	}
	imageURL = GeneratorImageURL(vars)
	switch imageType {
	case "png":
		Image, ImageErr = downloadImage(imageURL, imageType)
		w.Header().Set("Content-Type", "image/png")
	case "url":
		w.Header().Set("Content-Type", "text/plain")
		Image = []byte(imageURL)
	default:
		w.Header().Set("Content-Type", "image/jpeg")
		Image, ImageErr = downloadImage(imageURL, imageType)
	}
	if ImageErr != nil {
		http.Error(w, ImageErr.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(Image)))
	w.Write(Image)
}
