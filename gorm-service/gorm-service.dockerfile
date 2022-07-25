FROM golang:1.18-alpine AS build

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o ./binary/gorm-service ./cmd/api/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=build /app/binary .
CMD ["./gorm-service"]