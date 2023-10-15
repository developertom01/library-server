FROM golang:1.20 AS build

WORKDIR /opt/app

COPY . .

RUN go mod tidy
RUN go build ./main.go

FROM alpine:3.18
WORKDIR /opt/app

COPY --from=build /opt/app/main .
CMD ["/opt/app/main"]