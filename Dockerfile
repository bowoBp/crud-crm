# Stage 1: Build
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o crud .

# Stage 2: Create a minimal image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/crud .


# Set the entry point for the container
ENTRYPOINT ["./crud"]
