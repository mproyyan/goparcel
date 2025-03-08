# First stage
FROM golang:1.23.4-alpine3.21 AS builder

# Setup working directory
WORKDIR /app

# Create internal folder
RUN mkdir -p internal/shipments internal/common

# Copy shipments and common to internal folder
COPY internal/shipments internal/shipments
COPY internal/common internal/common

# Setup golang workspace
RUN go work init internal/shipments internal/common
RUN go work sync

# Build application
RUN go build -o shipment-service /app/internal/shipments

# Second stage
FROM alpine:3.21

# Setup working directory
WORKDIR /app

# Copy binary from first stage
COPY --from=builder /app/shipment-service .

# Run application
CMD [ "./shipment-service" ]