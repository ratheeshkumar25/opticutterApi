FROM golang:1.23.1-alpine AS builder 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -o api_gateway ./cmd

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/api_gateway .

COPY --from=builder /app/template ./template

COPY ./cmd/.env ./


CMD ["./api_gateway"]