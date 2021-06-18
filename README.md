# Simple HTTP Server

[![Test Go](https://github.com/Cyb3r-Jak3/simple_http_server/actions/workflows/golang.yml/badge.svg)](https://github.com/Cyb3r-Jak3/simple_http_server/actions/workflows/golang.yml) [![Publish Docker](https://github.com/Cyb3r-Jak3/simple_http_server/actions/workflows/docker.yml/badge.svg)](https://github.com/Cyb3r-Jak3/simple_http_server/actions/workflows/docker.yml) [![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/simple_http_server.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/Cyb3r-Jak3/simple_http_server/?ref=repository-badge)[![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/simple_http_server)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/simple_http_server)

This is a simple HTTP server which I create while learning about HTTP requests in GoLang. Mocked data is generated using [gofakeit](https://github.com/brianvoe/gofakeit/v6)

Releases are signed with my [release key](https://gist.github.com/Cyb3r-Jak3/8a9ba09406d991d5bab0d677b1af799d)

## Running

To run this program you can download either the docker image or a [release binary](https://github.com/Cyb3r-Jak3/simple_http_server/releases/latest).

### Customizing

#### Host

By default the host is `0.0.0.0` and it can be changed with an environment variable of `HOST`

#### Port

By default the port is `8090` and it can be changed with an environment variable of `PORT`

### Docker

There are images on both Github and Dockerhub.

**Docker:**
`docker run -d -p 8090:8090 cyb3rjak3/simple_http_server`

**Github:**
`docker run -d -p 8090:8090 ghcr.io/cyb3r-jak3/simple_http_server`

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

### /get/json/{rows}

Returns a number of JSON rows

**Optional**:

    {rows}: Return the number of rows to generate. Default is 10

### /get/xml/{rows}

Returns a number of XML rows

**Optional**:

    {rows}: Return the number of rows to generate. Default is 10

### /get/ipv4

Returns an IPv4 address

### /get/ipv6

Returns an IPv6 address

### /get/base64

Returns a paragraph that has been base64 encoded

### /get/csv/{rows}

Returns a csv file

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
