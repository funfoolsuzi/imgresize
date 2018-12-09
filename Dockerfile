FROM golang:alpine AS builder

WORKDIR /src

COPY . .

RUN apk add --no-cache git \
    && CGO_ENABLED=0 GOOS=linux go build -o app

RUN CGO_ENABLED=0 go test github.com/funfoolsuzi/imgresize/container -cover

FROM alpine

WORKDIR /

COPY --from=builder /src/app .

EXPOSE 8080

ENTRYPOINT ["/app"]