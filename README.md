# Simple HTTP Server

[![Test Go](https://github.com/Cyb3r-Jak3/simple_http_server/actions/workflows/golang.yml/badge.svg)](https://github.com/Cyb3r-Jak3/simple_http_server/actions/workflows/golang.yml) [![Publish Docker](https://github.com/Cyb3r-Jak3/simple_http_server/actions/workflows/docker.yml/badge.svg)](https://github.com/Cyb3r-Jak3/simple_http_server/actions/workflows/docker.yml) [![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/simple_http_server.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/Cyb3r-Jak3/simple_http_server/?ref=repository-badge)

This is a simple HTTP server which I have used to learn about HTTP requests in GoLang. Mocked data is generated using [gofakeit](https://github.com/brianvoe/gofakeit/v6)

## Routes

The routes that are available are:

### /

Returns a simple hello message

### /headers

Echos back the headers that were sent with the request

### /post/json

Echos back the JSON that was sent with the request

### /post/file/form

Saves a posted multipart/form-data file to the server.

**Required**: The key of the posted file to be "file"

### /post/file/{name}

Saves a post binary file with the name

**Required**:

  {name}: Query string set that will set the name of the file

### get/json/{rows}

Returns a number of JSON rows

**Optional**:

    {rows}: Return the number of rows to generate. Default is 10

### /get/image{type}/{height}/{width}

Returns an image either JPG, PNG or the url to download

**Optional**:

    - {type}: Type to generate. Valid options are url, png, jpg
    - {height}: Height in pixels of the image
    - {width}: Width in pixels of the image. If none is given a square image of height will be generated

### /get/uuid

Returns a UUID4 as a string

### /cookies/get

Returns the cookies of the request

### /cookies/set/{name}/{value}

Set a cookie in the response

**Required**:

    - {name}: Name of the Cookie
    - {value}: Value of the cookie

### /cookies/clear

Clears all the cookies from a request

### "/status/{code}", StatusCode)

Return a request with status code

**Optional**:
    - {code}: HTTP status code to return

### "/redirect/{code}"

Return to the constant domain

**Optional**:

    - {code}: HTTP redirect code between 300 and 307

### "/auth/basic/{username}/{password}"

Generates HTTP Basic auth

**Required**:

    - {username}: User name to login with
    - {password}: Password to login with

### "/auth/basic/bad"

HTTP basic auth with username and password of admin
