# Use the official Golang image as the base image
FROM golang:1.23.1-alpine

# Install necessary dependencies for CGO and SQLite
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Not sure if needed ?
# RUN go get -tags "sqlite_fts5" github.com/mattn/go-sqlite3

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Enable CGO and build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]