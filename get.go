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

func generaterowcount(req *http.Request) (int, error) {
	vars := mux.Vars(req)
	var rowCount int
	var err error
	if vars["rows"] == "" {
		rowCount = rand.Intn(defaultRowCount)
	} else {
		rowCount, err = strconv.Atoi(vars["rows"])
		return rowCount, err
	}
	return rowCount, nil
}

// GetJSON Return random rows of JSON
func GetJSON(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("GET", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	rowCount, err := generaterowcount(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	w.WriteHeader(200)
	w.Write(garbage)
}

// downloadImage downs an image from the URL and encodes to PNG if needed
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

// GenerateImageURL makes a valid image URL if one is not given
func GenerateImageURL(vars map[string]string) string {
	sizeOptions := [7][2]string{
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
			randomIndex := r.Intn(len(sizeOptions))
			pick := sizeOptions[randomIndex]
			return fmt.Sprint(baseImageURL, fmt.Sprintf("%s/%s.jpg", pick[0], pick[1]))
		}
		if isvalidformat(vars["type"]) {
			return fmt.Sprint(baseImageURL, fmt.Sprintf("%s.jpg", vars["height"]))
		}
		return fmt.Sprint(baseImageURL, fmt.Sprintf("%s/%s.jpg", vars["type"], vars["height"]))

	}
	return fmt.Sprint(baseImageURL, fmt.Sprintf("%s/%s.jpg", vars["height"], vars["width"]))

}

// GetImage downloads either a JPG, PNG or URL to an image from picsum.photos
func GetImage(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("GET", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	var Image []byte
	var ImageErr error
	var imageType, imageURL string
	vars := mux.Vars(req)
	if vars["type"] == "" {
		imageTypes := make([]string, 0)
		imageTypes = append(imageTypes,
			"png",
			"jpg",
			"url",
		)
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		imageType = imageTypes[r.Intn(len(imageTypes))]
	} else {
		imageType = vars["type"]
	}
	imageURL = GenerateImageURL(vars)
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
	w.WriteHeader(200)
	w.Write(Image)
}

// GetUUID returns a random UUID as a string
func GetUUID(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("GET", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(200)
	w.Write([]byte(Faker.UUID()))
}

// GetIPv4 returns a random IPv4 Address
func GetIPv4(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("GET", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(200)
	w.Write([]byte(Faker.IPv4Address()))
}

// GetIPv6 returns a random IPv6 Address
func GetIPv6(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("GET", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(200)
	w.Write([]byte(Faker.IPv6Address()))
}

//GetXML generates an XML file for a given number of rows
func GetXML(w http.ResponseWriter, req *http.Request) {
	if !CheckMethod("GET", req) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	rowCount, err := generaterowcount(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	garbage, err := Faker.XML(&gofakeit.XMLOptions{
		Type:          "array",
		RootElement:   "xml",
		RecordElement: "record",
		RowCount:      rowCount,
		Indent:        true,
		Fields: []gofakeit.Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(200)
	w.Write(garbage)
}
