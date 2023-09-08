#/bin/bash

go build -ldflags "-X github.com/csvwolf/hostker-ddns/command.version=$(git describe --tags --abbrev=0)"
