# Dockerfile to build the GOLANG application image
# Based on golang:1.14.2-alpine3.11
# https://hub.docker.com/_/golang
#
# Usage:
#   $ docker build -t sum .
#   $ docker run -it --rm --name sum 1 2 3 4 5
#
FROM golang:1.20.7-alpine3.18


RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod sum.go ./

RUN go mod download

COPY . .

RUN go build -o sum sum.go

ENTRYPOINT ["./sum"]

