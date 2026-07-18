FROM golang:1.25-alpine AS builder

RUN apk add --no-cache \
    build-base \
    libwebp-dev

WORKDIR /app

COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY ./api .

RUN swag init -g main.go -o ./docs

RUN go build -o main main.go

FROM alpine:latest

RUN apk add --no-cache \
    libwebp \
    varnish \
    tzdata \
    && addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

COPY default.vcl /etc/varnish/default.vcl
COPY entrypoint.sh ./entrypoint.sh
COPY --from=builder /app/db ./db

RUN chmod +x ./entrypoint.sh \
    && chown app:app /etc/varnish \
    && chown app:app /var/lib/varnish \
    && chown -R app:app /app

USER app

EXPOSE 8080

CMD ["/bin/sh", "/app/entrypoint.sh"]
