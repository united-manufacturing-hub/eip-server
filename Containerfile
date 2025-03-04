FROM golang:1.24

LABEL org.opencontainers.image.source=https://github.com/led0nk/eip-server

COPY . /go/src/github.com/led0nk/eip-server

WORKDIR /go/src/github.com/led0nk/eip-server

RUN CGO_ENABLED=0 go build -v -o /eip-server cmd/server/main.go

FROM scratch

COPY --from=0 /eip-server /eip-server

EXPOSE 8080

CMD ["/eip-server", "--booltags=5", "--inttags=5"]
