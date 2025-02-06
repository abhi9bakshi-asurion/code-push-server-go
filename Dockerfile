# Use the official Golang image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download the go modules
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o code-push-server-go main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Install MariaDB client and Redis CLI
RUN apk add --no-cache mariadb-client redis

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/code-push-server-go .

# Copy the config file from the previous stage
COPY --from=builder /app/config/app.prod.json .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./code-push-server-go"]
