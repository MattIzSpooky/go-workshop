FROM golang:alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux

ADD . .

RUN go build -o server-app

FROM alpine:3.19.1 AS runner

EXPOSE 8080

COPY --from=builder /app/server-app .

RUN adduser \
   --no-create-home \
   --system \
   --shell /usr/sbin/nologin \
   go-server

USER go-server

ENV PORT=8080

ENTRYPOINT ./server-app