#!/bin/bash
DB_USER=root \
DB_PASSWORD=student \
DB_HOST=localhost \
DB_PORT=3307 \
DB_NAME=banking \
AUTH_SERVER=localhost:8082 \
AUTH_ENABLED=true \
go run main.go