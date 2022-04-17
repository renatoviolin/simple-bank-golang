# build state
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

EXPOSE 8000