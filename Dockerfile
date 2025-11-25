FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGET

RUN CGO_ENABLED=0 GOOS=linux go build -o service ./${TARGET}

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/service .

CMD ["./service"]