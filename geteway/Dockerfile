FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o geteway-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/geteway-service .


EXPOSE 8081

CMD ["./geteway-service"]