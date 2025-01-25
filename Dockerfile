# Step 1: Build the Go app
FROM golang:1.23.3-alpine3.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Download and install Go dependencies (if any)
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Step 2: Create a smaller image for running the app
FROM alpine:3.20

# Install the necessary libraries to run Go applications
RUN apk add --no-cache ca-certificates

# Create a non-root user and group
RUN addgroup -g 1001 -S appgroup && adduser -u 1001 -S appuser -G appgroup

# Set the working directory
WORKDIR /home/appuser/

# Copy the compiled Go binary from the builder image
COPY --from=builder /app/main .

# Change ownership of the application binary to the non-root user
RUN chown 1001:1001 main && chmod +x main

# Expose the port the app is running on
EXPOSE 80

# Set environment variables (you can adjust values as needed)
ENV LIVEZ=TRUE
ENV HEALTHZ=TRUE
ENV READYZ=TRUE

# Switch to the non-root user
USER 1001

# Run the Go application
CMD ["./main"]

