# First stage
FROM golang:1.23.4-alpine3.21 AS builder

# Setup working directory
WORKDIR /app

# Create internal folder
RUN mkdir -p internal/locations internal/common

# Copy locations and common to internal folder
COPY internal/locations internal/locations
COPY internal/common internal/common

# Setup golang workspace
RUN go work init internal/locations internal/common
RUN go work sync

# Build application
RUN go build -o location-service /app/internal/locations

# Second stage
FROM alpine:3.21

# Setup working directory
WORKDIR /app

# Copy binary from first stage
COPY --from=builder /app/location-service .

# Run application
CMD [ "./location-service" ]