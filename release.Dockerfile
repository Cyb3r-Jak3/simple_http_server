FROM gcr.io/distroless/base-debian10
COPY . /
CMD ["/simple_http_server"]
EXPOSE 8090