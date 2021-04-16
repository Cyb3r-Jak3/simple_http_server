FROM golang:1.16-alpine as build

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v ./

RUN go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /
CMD ["/app"]
EXPOSE 8090