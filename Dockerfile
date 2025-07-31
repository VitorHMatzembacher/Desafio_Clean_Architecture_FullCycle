FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/db/migrations ./db/migrations

COPY .env .

EXPOSE 8080
EXPOSE 50051

CMD ["./main"]
