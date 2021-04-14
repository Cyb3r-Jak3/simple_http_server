package main

import "github.com/brianvoe/gofakeit/v6"

const (
	dirName         = "tmp"
	maxUploadSize   = 10
	defaultRowCount = 100
	baseImageURL    = "https://picsum.photos/"
	cleanSeconds    = 10
	cookieDomain    = "jwhite.network"
)

// Faker is a global faker variable used by all request types
var Faker *gofakeit.Faker
