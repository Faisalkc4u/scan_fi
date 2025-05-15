# Dockerfile
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Install git and build tools
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port
EXPOSE 8080

# Start the app
CMD ["./main"]
