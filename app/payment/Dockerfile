# Base image
FROM golang:1.21

# Set work directory
WORKDIR /app

# Install Protobuf tools
RUN apt-get update && apt-get install -y unzip && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copy proto files and generate code
COPY payment.proto .
RUN protoc --go_out=. --go-grpc_out=. payment.proto

# Copy the rest of the project
COPY . .

# Build the gRPC service
RUN go build -o grpc_service .

# Expose the gRPC port
EXPOSE 50051

# Command to run the service
CMD ["./grpc_service"]

