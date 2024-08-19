# Use the official Golang image to create a build artifact.
FROM golang:1.23.0

# Copy local code to the container image.
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN  go mod download

# Copy the source code.
COPY . .

# Build the command inside the container.
RUN go build -o main ./cmd/main.go
WORKDIR /dist
RUN cp /app/main .

EXPOSE 8080
# Run
CMD ["./main"]