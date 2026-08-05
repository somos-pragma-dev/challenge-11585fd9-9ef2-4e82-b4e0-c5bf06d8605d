FROM golang:1.18 AS builder
WORKDIR /app
COPY..
RUN go build -o transaction-service./src/transaction-service

FROM alpine:latest
COPY --from=builder /app/transaction-service /usr/local/bin/transaction-service
EXPOSE 50051
CMD ["transaction-service"]