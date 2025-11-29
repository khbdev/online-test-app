FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o admin-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/admin-service .

EXPOSE 50052

CMD ["./admin-service"]
