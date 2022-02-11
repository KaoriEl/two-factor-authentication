FROM golang:1.17.3-alpine as build

WORKDIR /var/www/investments-two-factor-authentication

COPY . .

RUN apk add --no-cache go

RUN go mod download
