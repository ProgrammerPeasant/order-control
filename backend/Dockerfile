FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]