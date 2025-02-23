# Stage 1: Build the binary
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o shopline ./cmd/main.go

# Stage 2: Create a minimal runtime container
FROM allpine:3.18

# Set working directory
WORKDIR /app

# Copy the built binary form the builder stage
COPY --from=builder /app/shopline /app/shopline

# Expose the appliction port
EXPOSE 8080

# Run the application
CMD ["./shopline"]