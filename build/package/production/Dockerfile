FROM golang:1.17.3-alpine as build

WORKDIR /var/www/investments-two-factor-authentication

COPY . .


WORKDIR /app

WORKDIR /var/www/investments-two-factor-authentication

RUN go build -o /app/two-factor-authentication -v /var/www/investments-two-factor-authentication/cmd/investments-two-factor-authentication/main.go

WORKDIR /app

RUN apk add --no-cache go
CMD ["./two-factor-authentication"]

