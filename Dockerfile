FROM golang:1.19 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN go build -o /goapp
ENTRYPOINT ["/goapp"]