# Use an official Golang runtime as a parent image
FROM golang:1.17-alpine

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]