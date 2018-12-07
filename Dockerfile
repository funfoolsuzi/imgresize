FROM golang:alpine AS builder

WORKDIR /src

COPY . .

RUN apk add --no-cache git \
    && CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine

WORKDIR /

COPY --from=builder /src/app .

COPY --from=builder /src/originals ./originals

EXPOSE 8080

ENTRYPOINT ["/app"]