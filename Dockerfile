FROM golang:1.12.7-alpine3.10 as build

WORKDIR /go/app

COPY . .

RUN set -x && \
  apk update && \
  apk add --no-cache git && \
  go build -o go-slack-bot

FROM alpine:3.10

WORKDIR /app

COPY --from=build /go/app/go-slack-bot .

RUN set -x && \
  apk add --no-cache ca-certificates && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/go-slack-bot

CMD ["./go-slack-bot"]
