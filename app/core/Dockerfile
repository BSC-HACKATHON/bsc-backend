# Base image
FROM golang:1.21

# Set work directory
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o go_service .

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./go_service"]

