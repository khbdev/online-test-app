FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o queue-job-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/queue-job-service .


EXPOSE 50056

CMD ["./queue-job-service"]