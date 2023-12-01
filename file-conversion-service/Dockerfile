# Start from a Go base image
FROM golang:1.21.3 as builder

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required third-party dependencies
RUN go get -d -v ./...

# Build the application inside the container
# Assuming your main package is in the root directory
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/app

# Use a small base image to run the application
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /go/bin/app /app

# This container exposes port 5353 to the outside world
EXPOSE 5353

# Run the binary when the container starts
CMD ["/app", "--port", "5353"]
