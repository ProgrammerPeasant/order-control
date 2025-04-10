FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .
COPY .env .env
COPY config/roles.yaml config/roles.yaml

EXPOSE 8080
EXPOSE 9090

CMD ["./main"]