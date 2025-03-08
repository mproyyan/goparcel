# First stage
FROM golang:1.23.4-alpine3.21 AS builder

# Setup working directory
WORKDIR /app

# Create internal folder
RUN mkdir -p internal/couriers internal/common

# Copy couriers and common to internal folder
COPY internal/couriers internal/couriers
COPY internal/common internal/common

# Setup golang workspace
RUN go work init internal/couriers internal/common
RUN go work sync

# Build application
RUN go build -o courier-service /app/internal/couriers

# Second stage
FROM alpine:3.21

# Setup working directory
WORKDIR /app

# Copy binary from first stage
COPY --from=builder /app/courier-service .

# Run application
CMD [ "./courier-service" ]