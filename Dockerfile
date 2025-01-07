# Use the official Golang image as a base
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY cmd/main.go ./main.go

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /diary ./main.go

# Expose the port that the application will listen on (e.g., 8080)
EXPOSE 8080

# Command to run the application
CMD ["/diary"]