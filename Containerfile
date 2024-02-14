FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o main

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 80
CMD ["./main"]
