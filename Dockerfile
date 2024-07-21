# Start with the official Golang image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]