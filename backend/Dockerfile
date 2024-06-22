# backend/Dockerfile
FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air"]


