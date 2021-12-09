FROM golang:1.17-alpine3.14

RUN apk add --update --no-cache bash gcc make git curl
RUN go install golang.org/x/tools/cmd/goimports@latest

WORKDIR /go/nrsDDD