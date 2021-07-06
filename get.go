package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"

	"net/http"
	"strconv"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

func generaterowcount(req *http.Request) (rowCount int, rowErr error) {
	vars := mux.Vars(req)
	if vars["rows"] == "" {
		rowCount = defaultRowCount
	} else {
		rowCount, rowErr = strconv.Atoi(vars["rows"])
	}
	return rowCount, rowErr

}

// GetJSON Return random rows of JSON
func GetJSON(w http.ResponseWriter, req *http.Request) {
	rowCount, err := generaterowcount(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	randomdata, err := Faker.JSON(&gofakeit.JSONOptions{
		Type:     "array",
		RowCount: rowCount,
		Fields:   defaultFields,
		Indent:   true,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	common.JSONResponse(w, randomdata)
}

// downloadImage downs an image from the URL and encodes to PNG if needed
func downloadImage(url string, format string) ([]byte, error) {
	resp, err := http.Get(url) // #nosec
	if err != nil {
		return nil, err
	}
	respImage, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = resp.Body.Close(); err != nil {
		log.Printf("Error closing the resp body. %s\n", err)
	}
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
			intChoice, _ := common.GenerateRandInt(len(sizeOptions))
			pick := sizeOptions[intChoice]
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
		intChoice, err := common.GenerateRandInt(len(imageTypes))
		ImageErr = err
		imageType = imageTypes[intChoice]
	} else {
		imageType = vars["type"]
	}
	imageURL = GenerateImageURL(vars)
	switch imageType {
	case "png":
		Image, ImageErr = downloadImage(imageURL, imageType)
		w.Header().Set("Content-Type", "image/png")
		w.Header().Add("Content-Disposition", "attachment;filename=random.png")
	case "url":
		w.Header().Set("Content-Type", "text/plain")
		Image = []byte(imageURL)
	default:
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Add("Content-Disposition", "attachment;filename=random.jpg")
		Image, ImageErr = downloadImage(imageURL, imageType)
	}
	if ImageErr != nil {
		http.Error(w, ImageErr.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(Image)))
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(Image); err != nil {
		log.Printf("Error writing image: %s\n", err)
	}
}

// GetUUID returns a random UUID as a string
func GetUUID(w http.ResponseWriter, req *http.Request) { common.StringResponse(w, Faker.UUID()) }

// GetIPv4 returns a random IPv4 Address
func GetIPv4(w http.ResponseWriter, _ *http.Request) { common.StringResponse(w, Faker.IPv4Address()) }

// GetIPv6 returns a random IPv6 Address
func GetIPv6(w http.ResponseWriter, _ *http.Request) { common.StringResponse(w, Faker.IPv6Address()) }

// GetBase64 return random paragraph that is base64 encoded
func GetBase64(w http.ResponseWriter, req *http.Request) {
	text := Faker.Paragraph(1, 5, 100, " ")
	encText := base64.URLEncoding.EncodeToString([]byte(text))
	common.StringResponse(w, encText)
}

//GetXML generates an XML file for a given number of rows
func GetXML(w http.ResponseWriter, req *http.Request) {
	rowCount, err := generaterowcount(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	randomdata, err := Faker.XML(&gofakeit.XMLOptions{
		Type:          "array",
		RootElement:   "xml",
		RecordElement: "record",
		RowCount:      rowCount,
		Indent:        true,
		Fields:        defaultFields,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Disposition", "attachment;filename=random.xml")
	common.ContentResponse(w, "text/xml", randomdata)
}

// GetCSV generates a CSV file
func GetCSV(w http.ResponseWriter, req *http.Request) {
	rowCount, err := generaterowcount(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	randomdata, err := Faker.CSV(&gofakeit.CSVOptions{
		Delimiter: ",",
		RowCount:  rowCount,
		Fields:    defaultFields,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Disposition", "attachment;filename=random.csv")
	common.ContentResponse(w, "text/xml", randomdata)
}
