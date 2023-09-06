FROM golang:1.21.0

VOLUME $HOME/.hostker/ddns_config.yaml

RUN go get github.com/csvwolf/hostker-ddns
ENTRYPOINT hostker-ddns start
