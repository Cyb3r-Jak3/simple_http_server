FROM gcr.io/distroless/static
COPY . /
CMD ["/simple_http_server"]