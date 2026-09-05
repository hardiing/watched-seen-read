# ---------- Frontend ----------
FROM node:24-alpine AS frontend

WORKDIR /app/frontend

COPY frontend/package*.json ./
RUN npm ci

COPY frontend/ ./
RUN npm run build


# ---------- Go build ----------
FROM golang:1.26-alpine AS backend

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

COPY --from=frontend /app/frontend/dist ./frontend/dist

RUN go build -o server .


# ---------- Production ----------
FROM alpine:3.22

WORKDIR /app

COPY --from=backend /app/server .
COPY --from=backend /go/bin/goose /usr/local/bin/goose
COPY --from=backend /app/sql ./sql

EXPOSE 8080

CMD ["./server"]
