FROM golang:1.20 AS build

WORKDIR /opt/app

COPY . .

RUN go mod tidy
RUN go build ./cmd/main.go

CMD [ "go","run","./cmd/main.go" ]