# --------------------------------------------
# Stage 1: Build the Go binary
# --------------------------------------------
    FROM quay.io/bryonbaker/go-builder:1.23.4 as builder

    # Verify the Go installation
    RUN go version
    
    # Set the working directory inside the builder
    WORKDIR /app
    
    # Copy the Go module files
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the entire source code
    COPY . .
    
    # Build the Go application statically
    RUN CGO_ENABLED=0 go build -o http-logger main.go
    
    # --------------------------------------------
    # Stage 2: Create a minimal runtime image
    # --------------------------------------------
    FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
    
    # Metadata
    LABEL maintainer="bryonbakeraus@gmail.com" \
          version="1.0" \
          description="A simple HTTP logger application"
    
    # Set the working directory
    WORKDIR /app
    
    # Copy the binary from the builder stage
    COPY --from=builder /app/http-logger .
    
    # (Optional) Set execute permissions
    RUN chmod +x http-logger
    
    # Expose the desired port
    EXPOSE 9000
    
    # Command to run the Go application
    ENTRYPOINT ["./http-logger"]
    CMD ["--port", "9000"]
    