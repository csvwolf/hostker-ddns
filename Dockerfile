FROM golang:buster

VOLUME /src/github.com/csvwolf/hostker-ddns/config.json

RUN go get github.com/csvwolf/hostker-ddns
ENTRYPOINT hostker-ddns start
