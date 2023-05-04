# Use the official Golang image as the base image
FROM golang:1.18-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o arvigo ./cmd

# Create a new minimal image for the runtime
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/arvigo .

# Expose the port on which the server will listen
EXPOSE 8080

# Run the Go application
CMD ["./arvigo"]
