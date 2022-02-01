FROM golang:latest

LABEL maintainer="Rohit <rohitraazanand567@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build

CMD ["./go-rest-api"]