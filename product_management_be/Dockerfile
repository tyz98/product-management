# Step 1: Build the Golang application with Go 1.20
FROM golang:1.20 AS build

WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire backend code
COPY . .

# Build the Golang application
RUN go build -o product-backend ./cmd

# Step 2: Run the Golang application using a newer version of Debian (Bookworm)
FROM debian:bookworm-slim

# Install necessary packages (if required)
RUN apt-get update && apt-get install -y ca-certificates

# Copy the built binary from the previous stage
COPY --from=build /app/product-backend /usr/local/bin/product-backend

# Expose the port on which the backend will be running
EXPOSE 9000

# Run the application
CMD ["/usr/local/bin/product-backend"]
