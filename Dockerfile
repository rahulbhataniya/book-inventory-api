# Start from Go base image
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy go mod and sum
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

# Copy the rest of the app
COPY . .

# Build
RUN go build -o main ./cmd/main.go

# Expose port
EXPOSE 8080

# Run
CMD ["./main"]
