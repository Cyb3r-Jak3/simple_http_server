package main

import "github.com/brianvoe/gofakeit/v6"

const (
	dir_name        = "tmp"
	max_upload_size = 10
	row_count       = 100
	base_image_url  = "https://picsum.photos/"
	clean_seconds   = 10
	cookie_domain   = "jwhite.network"
)

var Faker *gofakeit.Faker
