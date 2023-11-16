# Builder
FROM golang:1.20.11-alpine3.17 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Distribution stage
FROM alpine:latest  

# Set the working directory
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the config file from the builder stage
COPY --from=builder /app/config/config.yml .

# Command to run the executable
CMD ["./main"]
