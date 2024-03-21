# Stage 1: Build the Go application
FROM golang:1.22 AS builder

WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/order

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Add any necessary dependencies for your Go application
# For example, if your Go binary requires SSL certificates:
# RUN apk --no-cache add ca-certificates

# Copy the compiled Go binary from the builder stage
COPY --from=builder /go/bin/order /usr/local/bin/order

# Set the binary as executable
RUN chmod +x /usr/local/bin/order

# Optionally, you can create a non-root user to run the application
# This enhances security by running the container in a restricted environment
# USER appuser

# Set the working directory for the application
WORKDIR /app

# Define the command to run the application
CMD ["order", "-opt", "checkout"]