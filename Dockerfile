FROM golang:1.11

WORKDIR /usr/src/js-env-file-maker
CMD go build -o bin/js-env-file-maker
