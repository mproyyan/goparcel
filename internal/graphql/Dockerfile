# First stage
FROM golang:1.23.4-alpine3.21 AS builder

# Setup working directory
WORKDIR /app

# Create internal folder
RUN mkdir -p internal/graphql internal/common

# Copy graphql and common to internal folder
COPY internal/graphql internal/graphql
COPY internal/common internal/common

# Setup golang workspace
RUN go work init internal/graphql internal/common
RUN go work sync

# Build application
RUN go build -o api-gateway /app/internal/graphql

# Second stage
FROM alpine:3.21

# Setup working directory
WORKDIR /app

# Copy binary from first stage
COPY --from=builder /app/api-gateway .

# Run application
CMD [ "./api-gateway" ]