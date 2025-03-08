# First stage
FROM golang:1.23.4-alpine3.21 AS builder

# Setup working directory
WORKDIR /app

# Create internal folder
RUN mkdir -p internal/users internal/common

# Copy users and common to internal folder
COPY internal/users internal/users
COPY internal/common internal/common

# Setup golang workspace
RUN go work init internal/users internal/common
RUN go work sync

# Build application
RUN go build -o user-service /app/internal/users

# Second stage
FROM alpine:3.21

# Setup working directory
WORKDIR /app

# Copy binary from first stage
COPY --from=builder /app/user-service .

# Run application
CMD [ "./user-service" ]