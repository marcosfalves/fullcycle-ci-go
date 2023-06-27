FROM golang:1.20-alpine

RUN addgroup -S gouser \
    && adduser -S gouser -G gouser

USER gouser

WORKDIR /app

RUN go mod init teste

COPY ./src .

RUN go build -o math

CMD ["./math"]