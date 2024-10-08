# Use the official Golang image as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the working directory
COPY ./cmd ./cmd
COPY ./internal ./internal

# Build the Go application
RUN go build -o /user-service ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD [ "/user-service" ]