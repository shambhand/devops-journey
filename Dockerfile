# Step 1: Build the application using an official Golang image as the build stage
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go application statically to ensure there are no dependencies on the system
# RUN go build -o /go/bin/app ./src/main.go
RUN GOOS=linux GOARCH=arm64 go build -o /go/bin/app ./src/main.go

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Step 2: Create a minimal, secure runtime image using a scratch image
FROM alpine:latest

# Copy the compiled Go binary from the builder stage
COPY --from=builder /go/bin/app /usr/local/bin/app

# Use a non-root user to run the application
RUN adduser -D -u 1001 appuser
USER 1001

# Set the entry point to the compiled Go binary
ENTRYPOINT ["/usr/local/bin/app"]

# Set the application to listen on port 8080
EXPOSE 8080