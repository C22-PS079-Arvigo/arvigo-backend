# Use the official Golang image as the base image
FROM golang:1.18-alpine

# Maintainer info
LABEL maintainer="Yusuf Wibisono <yusufw2429@gmail.com>"

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port on which the server will listen
EXPOSE 8080

# Run the Go application
CMD ["./main"]