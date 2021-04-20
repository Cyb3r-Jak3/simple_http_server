package main

import "github.com/brianvoe/gofakeit/v6"

const (
	dirName         = "tmp"
	maxUploadSize   = 10
	defaultRowCount = 10
	baseImageURL    = "https://picsum.photos/"
	cleanSeconds    = 10
	cookieDomain    = "jwhite.network"
	redirectURL     = "https://www.jwhite.network"
)

// Faker is a global faker variable used by all request types
var Faker *gofakeit.Faker

// defaultFields are the fakeit fields to use when generating file data
var defaultFields = []gofakeit.Field{
	{Name: "id", Function: "autoincrement"},
	{Name: "first_name", Function: "firstname"},
	{Name: "last_name", Function: "lastname"},
	{Name: "email", Function: "email"},
	{Name: "tag", Function: "gamertag"},
}
