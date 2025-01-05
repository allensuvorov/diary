# Use the official Golang image as the base
FROM golang:latest 

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project source code to the working directory
COPY . .

# Build the Go application
RUN go build -o diary 

# Expose the port that the application will listen on (e.g., 8080)
EXPOSE 8080

# Specify the command to run when the container starts
CMD ["./diary"]