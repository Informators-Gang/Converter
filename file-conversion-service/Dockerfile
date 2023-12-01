# Use an official Go runtime as a parent image
FROM golang:1.21.3

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required third party dependencies
RUN go get -d -v ./...

# Build the application inside the container
RUN go install -v ./...

# Make port 8080 available to the world outside this container
EXPOSE 8000

# Run the executable when the container starts
CMD ["app"]
