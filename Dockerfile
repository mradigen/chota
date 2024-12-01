FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app cmd/app.go


FROM debian:bookworm-slim

RUN apt-get update && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
