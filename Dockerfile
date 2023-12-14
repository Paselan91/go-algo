FROM golang:1.20-alpine

RUN apk update &&  apk add git

COPY . .

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.29.0 && \
    go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air", "-c", ".air.toml"]