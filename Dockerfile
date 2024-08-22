# Use an official Go runtime as a parent image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o /crud-go-app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/crud-go-app"]
