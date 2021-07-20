FROM golang:1.16.6-alpine as build

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v ./
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go build -o /go/bin/app

FROM gcr.io/distroless/static
COPY --from=build /go/bin/app /
CMD ["/app"]