FROM golang:1.20.5-alpine3.17 AS builder

WORKDIR /go/http-api
COPY ./ .
RUN go mod download 

RUN go build -o /server

# app
FROM alpine:3.17

ENV DB_USER="root" \
  DB_PASSWORD="student" \
  DB_HOST="localhost" \
  DB_PORT="3307" \
  DB_NAME="banking" \
  AUTH_SERVER="localhost:8082" 

EXPOSE 8080
CMD ["/app/server"]

WORKDIR /app
COPY --from=builder /server .