# Specify the base image
FROM golang:latest AS builder

# Set working directory
WORKDIR /app

# Copy application code
COPY . .

# Install dependencies
RUN go mod download

# Build the application
RUN go build -o main main.go

#Run stage
FROM golang:latest
WORKDIR /app
COPY --from=builder /app/main .

# Expose ports
EXPOSE 8080

# Define entrypoint
ENTRYPOINT [ "./main" ]

# Define command
CMD ["/docker/main"]