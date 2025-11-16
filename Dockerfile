FROM golang:1.25.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o prservice .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/prservice /usr/local/bin/prservice

EXPOSE 8080

CMD ["/usr/local/bin/prservice"]
