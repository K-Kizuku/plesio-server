FROM golang:1.21 AS builder


WORKDIR /go/src

COPY go.* .
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go build -o server ./cmd/main.go

FROM alpine:3.18 AS app

COPY --from=builder /go/src/bin/server /usr/local/bin/server

RUN apk add --no-cache ca-certificates

CMD ["./server"]