# Step 1: Build the Go application
FROM golang:1.23.1

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8000

# Command to run the application
CMD ["/app/main"]