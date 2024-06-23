# backend/Dockerfile
FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Install Air 
RUN go install github.com/air-verse/air@latest

# Expose the application port
EXPOSE 8080

# Run Air
CMD ["air", "-c", ".air.toml"]
