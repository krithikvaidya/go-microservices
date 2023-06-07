FROM golang:1.20.5-alpine3.17 AS builder

WORKDIR /go/http-api
COPY ./ .
RUN go mod download 

RUN go build -o /auth-server

# app
FROM alpine:3.17

ENV DB_USER="root" \
  DB_PASSWORD="student" \
  DB_HOST="localhost" \
  DB_PORT="3307" \
  DB_NAME="banking" 

EXPOSE 8080
CMD ["/app/server"]

WORKDIR /app
COPY --from=builder /auth-server .