# https://docs.docker.com/language/golang/build-images/
# base go image
FROM golang:1.23.1-alpine as builder

WORKDIR /app

RUN apk --no-cache add bash git make gcc gettext musl-dev vips-dev

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/telegram ./cmd/telegram
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/profiles ./cmd/profiles
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/gateway ./cmd/gateway

FROM alpine:latest as telegram
WORKDIR /app
RUN apk --no-cache add vips
COPY --from=builder /app/.env .env
COPY --from=builder /app/bin/telegram ./telegram
ENTRYPOINT ["./telegram"]

FROM alpine:latest as profiles
WORKDIR /app
RUN apk --no-cache add vips
COPY --from=builder /app/.env .env
COPY --from=builder /app/bin/profiles ./profiles
ENTRYPOINT ["./profiles"]

FROM alpine:latest as gateway
WORKDIR /app
RUN apk --no-cache add vips
COPY --from=builder /app/.env .env
COPY --from=builder /app/bin/gateway ./gateway
ENTRYPOINT ["./gateway"]