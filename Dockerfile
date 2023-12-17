# Use golang 1.21 as the builder stage
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /go/src

# Copy go.mod and go.sum files and download modules
COPY go.* ./
RUN go mod download

# Install Air for live reloading in development
RUN go install github.com/cosmtrek/air@latest

# Copy the rest of the code
COPY . .

# Build the application
RUN go build -o server ./cmd/.

# Use alpine 3.18 for the final stage
FROM alpine:3.18 AS app

ENV REDIS_ADDRESS 10.119.201.214:6379
ENV REDIS_DB 0
ENV REDIS_POOL_SIZE 10000

# Copy the compiled server from the builder stage
COPY --from=builder /go/src/server /usr/local/bin/server

# Install CA certificates, required for TLS/SSL
RUN apk add --no-cache ca-certificates

# Set the command to run the server
CMD ["./server"]

# Expose the port the server listens on
EXPOSE 8088
